package config

//TODO: Move Viper to here and common file? @Michel

type textFormat struct {
	Color string
	Reset string
}

type Configurations interface {
	SetTextFormatting(colour string) *textFormat
}

func NewTextConfig() Configurations {
	return &textFormat{}
}

func (*textFormat) SetTextFormatting(colour string) *textFormat {

	//default is set to white
	tf := textFormat{
		Color: "\033[37m",
		Reset: "\033[0m",
	}

	switch colour {
	case "red":
		tf.Color = "\033[31m"
	case "green":
		tf.Color = "\033[32m"
	case "yellow":
		tf.Color = "\033[33m"
	case "blue":
		tf.Color = "\033[34m"
	case "purple":
		tf.Color = "\033[35m"
	case "cyan":
		tf.Color = "\033[36m"
	case "white":
		tf.Color = "\033[37m"
	default:
		tf.Color = "\033[37m"
	}

	return &tf
}
