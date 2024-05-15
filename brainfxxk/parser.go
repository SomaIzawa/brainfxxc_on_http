package brainfxxk

import (
	"fmt"
	"http_on_brainfxxk/util"
)

type Parser struct {
	Code []string   // BrainFxxkコードを1byte区切りで分割した文字列配列
	CIndex int      // コードの実行位置
	Memory []int    // 記憶領域
	MIndex int      // 記憶領域のポインタ
	LoopStart []int // ループの開始位置の記憶領域
	LSIndex int     // ループの開始位置の記憶領域のポインタ
	OutputString string
	MaxStep int
	StepCount int
}

// コンストラクタ

func NewParser(code string, memorySize int, maxStep int) Parser{
	return Parser{
		Code: util.ExtractSpecificCharacters(code, []string{"+","-",">","<","[","]",".",","}),
		CIndex: 0,
		Memory: make([]int, memorySize),
		MIndex: memorySize / 2,
		LoopStart: make([]int, len(code)/2),
		LSIndex: -1,
		OutputString: "",
		MaxStep: maxStep,
		StepCount: 0,
	}
}

// 実行

func (p *Parser) Exec() error {
	for p.CIndex < len(p.Code) {
		//評価
		if err := p.EvaluateCode(p.Code[p.CIndex]); err != nil {
			return err
		}
		p.StepCount++
		if p.MaxStep < p.StepCount {
			return fmt.Errorf("RuntimeError: maximum step count reached (the process is taking too long)")
		}
	}
	return nil
}

// 命令評価

func (p *Parser) EvaluateCode(str string) error {
	var err error
	switch str {
	case "+":
		err = p.MCountUp()
	case "-":
		err = p.MCountDown()
	case ">":
		err = p.MPInc()
	case "<":
		err = p.MPDec()
	case "[":
		err = p.ProcessLoopStart()
	case "]":
		err = p.ProcessLoopEnd()
	case ".":
		err = p.Output()
	case ",":
		err = p.Input()
	}
	return err
}

// ステップを次へ

func (p *Parser) next() {
	p.CIndex++
}

// 各命令に対応する処理

func (p *Parser) MCountUp() error {
	p.Memory[p.MIndex]++
	p.next()
	return nil
}

func (p *Parser) MCountDown() error {
	p.Memory[p.MIndex]--
	p.next()
	return nil
}

func (p *Parser) MPInc() error {
	if p.MIndex != (len(p.Memory) - 1) {
		p.MIndex++
	} else {
		return fmt.Errorf("RuntimeError: Failed to increment the memory pointer")
	}
	p.next()
	return nil
}

func (p *Parser) MPDec() error {
	if p.MIndex != 0 {
		p.MIndex--
	} else {
		return fmt.Errorf("RuntimeError: Failed to decrement the memory pointer")
	}
	p.next()
	return nil
}

func (p *Parser) ProcessLoopStart() error  {
	p.LSIndex++
	p.LoopStart[p.LSIndex] = p.CIndex
	p.next()
	return nil
}

func (p *Parser) ProcessLoopEnd() error {
	if p.Memory[p.MIndex] > 0 {
		p.CIndex = p.LoopStart[p.LSIndex]
		p.next()
	} else {
		p.LoopStart[p.LSIndex] = 0
		p.LSIndex--
		p.next()
	}
	return nil
}

func (p *Parser) Output() error {
	s := fmt.Sprintf("%c", p.Memory[p.MIndex])
	fmt.Printf("%s", s)
	p.OutputString = p.OutputString + s
	p.next()
	return nil
}

func (p *Parser) Input() error {
	var str string
  fmt.Scan(&str)
	firstChar := str[0]
	p.Memory[p.MIndex] = int(firstChar)
	p.next()
	return nil
}

// 以下、ディベロッパー用

func (p *Parser) ShowMemory(width int){
	var head []int
	var output []int
	outputWidth := 0
	height := len(p.Memory) / width + 1
	fmt.Print("\n\n")
	println("Memory:")
	println("")
	for i, item := range p.Memory {
		if i % width == 0 && i / width + 1 < height {
			outputWidth = width
			util.OutPutLine(outputWidth)
		} else if i % width == 0 && i / width + 1 >= height {
			outputWidth = len(p.Memory) % width
			util.OutPutLine(outputWidth)
		}

		head = append(head, i)
		output = append(output, item)

		if (i % width == width - 1 && i / width + 1 < height) ||
		((i  % width) == (len(p.Memory) % width - 1) && i / width + 1 >= height) {
			util.OutPutValues(head)
			util.OutPutLine(outputWidth)
			util.OutPutValues(output)
			util.OutPutLine(outputWidth)
			head = []int{}
			output = []int{}
			util.OutPutEmptyLine()
		}
	}
}