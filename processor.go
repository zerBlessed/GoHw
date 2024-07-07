// Необходимо имплементировать интерфейс Manager так, чтобы Manager мог принимать данные из одного Reader
// обрабатывать полученные данные на каждом из списка Processor и результирующие данные передавать в Writer.
// При возникновении ошибки при обработке, прочитанные из Reader данные необходимо пропустить.
package main

type Data struct {
	ID      int
	Payload map[string]interface{}
}

type Reader interface {
	Read() []*Data
}

type Processor interface {
	Process(Data) ([]*Data, error)
}

type Writer interface {
	Write([]*Data)
}

type Manager struct {
	reader     Reader
	processors []Processor
	writer     Writer
}

func (m *Manager) Manage() {
	for {
		data := m.reader.Read()
		if len(data) == 0 {
			continue
		}

		for _, processor := range m.processors {
			processedData := make([]*Data, 0)
			for _, d := range data {
				result, err := processor.Process(*d)
				if err != nil {
					continue
				}
				processedData = append(processedData, result...)
			}

			m.writer.Write(processedData)
		}
	}
}
