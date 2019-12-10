package intcode

func init() {
	for _, instruction := range Instructions {
		InstructionNames[instruction.Name()] = instruction
	}
}
