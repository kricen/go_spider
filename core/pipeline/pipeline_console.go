package pipeline

import (
	"github.com/kricen/go_spider/core/common/com_interfaces"
	"github.com/kricen/go_spider/core/common/page_items"
)

type PipelineConsole struct {
}

func NewPipelineConsole() *PipelineConsole {
	return &PipelineConsole{}
}

func (this *PipelineConsole) Process(items *page_items.PageItems, t com_interfaces.Task) {
}
