package main

import (
  "fmt"
  "time"
  "strings"
  "strconv"
)

type Animation struct {
  frames [][]string
  loop int
  delay time.Duration
  timeout chan bool
}

func CreateAnimation(str [][]string, frameNum int, t time.Duration) *Animation {
  if (len(str) == 0) {
    fmt.Printf("Animation needs animation frames")
  }

  s := &Animation{
    frames: str,
    loop: frameNum,
    delay: t,
    timeout: make(chan bool),
  }

  return s
}


func (a *Animation) Start() {
  go func(){
    ticker := time.NewTicker(a.delay)
    ticksPerSec := int(time.Second/a.delay)
    ticks := 0
    go func(){
      framePos := 0
      for range ticker.C {
        for i := range a.frames {
          if (len(a.frames[i]) <= framePos){
            fmt.Printf("\n"+strings.Replace(a.frames[i][len(a.frames[i])-1], "${time}", strconv.Itoa(ticks/ticksPerSec+1), -1))
          } else {
            fmt.Printf("\n\033[K"+strings.Replace(a.frames[i][framePos], "${time}", strconv.Itoa(ticks/ticksPerSec+1), -1))
          }
        }

        // erase all animation lines
        for range a.frames {
          fmt.Printf("\033[A\r") // cursor up + newline
        }

        if (framePos >= a.loop-1){
          framePos = 0
        } else {
          framePos = framePos + 1
        }
        ticks = ticks + 1
      }
    }()

    <-a.timeout
    ticker.Stop()
  }()
}

func (a *Animation) ChangeFrame(newFrame [][]string) {
  a.timeout <- true
  // wait till the current frame is done printing
  time.Sleep(10*time.Millisecond)

  // now clear the current frame
  for range a.frames {
    fmt.Printf("\033[B\033[K") // cursor down, delete line
  }
  for range a.frames {
    fmt.Printf("\033[A") // go back up
  }

  a.frames = newFrame // set up new frame
  a.Start() // start again
}

func main() {
  fmt.Printf("   7 MIN WORKOUT TIME!   ")
  frames := [][]string {
    {"!!! GET READY !!!"},
    {"┏ (-_-)┓","┏ (-_-)┛", "┗ (-_-)┓"},
  }
  a := CreateAnimation(frames, 3, 500*time.Millisecond)
  a.Start()
  time.Sleep(5 * time.Second)

  excercises := [][][]string {
    {
      {"JUMPING JACKS! ${time}"},
      {" \\(-_-)/ ","--(-_-)--", "┏ (-_-)┓ "},
    },
    {
      {"WALL SIT! ${time}"},
      {"༼ง=ಠ益ಠ=༽ง ","(ᕗ° ਊ ͠°)ᕗ ", "(ง ͡ʘ ͜ʖ ͡ʘ)ง "},
    },
    {
      {"PUSH UP! ${time}"},
      {" ᕦ( ͡° ͜ʖ ͡°)ᕤ "," ᕦʕ ° o ° ʔᕤ ", " ᕙ(◉෴◉)ᕗ "},
    },
    {
      {"CRUNCHES! ${time}"},
      {"༼; ಠ ਊ ಠ༽","(╬ ┛ಠ益ಠ)┛", "ᕙ( ͡° ͜ʖ ͡°)ᕗ"},
    },
    {
      {"CHAIR STEPS! ${time}"},
      {"     (ﾉಠдಠ)ﾉ︵┻━┻   ","      ┻━┻ ヘ╰( •̀ε•́ ╰)  ", "┻━┻ ︵ ¯\\_༼ᴼل͜ᴼ༽_/¯ ︵ ┻━┻"},
    },
    {
      {"SQUAT! ${time}"},
      {"୧| ͡ᵔ ﹏ ͡ᵔ |୨","ԅ[ ﹒︣ ͜ʟ ﹒︣ ]ﾉ", "へ། ¯͒ ʖ̯ ¯͒ །ᕤ"},
    },
    {
      {"TRICEP DIPS! ${time}"},
      {"┏༼ ◉ ╭╮ ◉༽┓","┌໒(: ⊘ ۝ ⊘ :)७┐", "〳: ⊘ ڡ ⊘ :〵"},
    },
    {
      {"PLANK! ${time}"},
      {"╏ ˵ ・ ͟ʖ ・ ˵ ╏","║ * ರ Ĺ̯ ರ * ║", "╚▒ᓀ▃ᓂ▒╝"},
    },
    {
      {"HIGH KNEES! ${time}"},
      {"⋋( ◕ ∧ ◕ )⋌","⋋╏ ᓀ 〜 ᓂ ╏⋌", "⋋( ” ͠° ʖ̫ °͠ ” )⋌"},
    },
    {
      {"LUNGE! ${time}"},
      {"┌| ◔ ▃ ◔ |┐","ᕕ║ ° ڡ ° ║┐", "╭(๑¯д¯๑)╮"},
    },
    {
      {"PUSH UP & ROTATION! ${time}"},
      {"へʕ ∗ ´ ۝ ´ ∗ ʔ┘","┌໒( : ͡° д °͡ : )७┐", "c༼ ” ͡° ▃ °͡ ” ༽ᕤ"},
    },
    {
      {"SIDEPLANK! ${time}"},
      {"‎(/.__.)/","ε(´סּ︵סּ`)з", "( ⋟﹏⋞ )"},
    },
    {
      {"YOU'RE DONE! PARTY TIME!"},
      {"╰(•̀ 3 •́)━☆ﾟ.*･｡ﾟ","ヽ༼ຈل͜ຈ༽⊃─☆*:・ﾟ", "༼∩ຈل͜ຈ༽つ━☆ﾟ.*･｡ﾟ", ". * ･ ｡ﾟ☆━੧༼ •́ ヮ •̀ ༽୨"},
    },
  }

  for i := range excercises {
    a.ChangeFrame(excercises[i])
    time.Sleep(30 * time.Second)
  }
}
