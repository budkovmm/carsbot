package pdfgen

import (
	"fmt"

	"carsbot/internal/state"

	"codeberg.org/go-pdf/fpdf"
)

// GenerateContractPDF генерирует договор купли-продажи и сохраняет PDF по указанному пути
func GenerateContractPDF(st *state.UserState, outPath string) error {
	pdf := fpdf.New("P", "mm", "A4", "")
	pdf.AddPage()

	// Добавляем поддержку кириллицы через UTF-8 шрифт
	fontPath := "DejaVuSans.ttf" // Файл шрифта должен быть в рабочей директории
	pdf.AddUTF8Font("DejaVu", "", fontPath)
	pdf.SetFont("DejaVu", "", 14)

	pdf.Cell(0, 10, "ДОГОВОР КУПЛИ-ПРОДАЖИ АВТОМОБИЛЯ")
	pdf.Ln(12)

	pdf.SetFont("DejaVu", "", 12)
	pdf.MultiCell(0, 8, fmt.Sprintf(
		"Город: %s    Дата: %s\n\nПродавец: %s\nПокупатель: %s\n\nАвтомобиль:\nVIN: %s\nМарка/Модель: %s\nГод выпуска: %s\nЦвет: %s\n\nСтоимость: %s руб.\n\nСтороны договорились о передаче автомобиля по указанной цене.\n\nПодписи сторон:\n\n__________________    __________________\nПродавец              Покупатель",
		st.City, st.Date, st.SellerName, st.BuyerName, st.VIN, st.BrandModel, st.Year, st.Color, st.Price,
	), "", "L", false)

	// Сохраняем PDF
	if err := pdf.OutputFileAndClose(outPath); err != nil {
		return fmt.Errorf("ошибка сохранения PDF: %w", err)
	}
	return nil
}
