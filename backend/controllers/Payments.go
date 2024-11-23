package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/skip2/go-qrcode"
)

func GenerateQR(c *gin.Context) {
	// รับจำนวนเงินจากคำขอ
	var request struct {
		Amount int `json:"amount"`
	}
	if err := c.BindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	// สร้าง PromptPay QR Code
	promptPayID := "0847866591" // หมายเลข PromptPay (ใส่ของร้านค้า)
	amount := float64(request.Amount)
	qrString := generatePromptPayQR(promptPayID, amount)

	// คำนวณ CRC16
	crcString := qrString + "6304" // ใส่ "6304" หลังข้อมูล QR
	crc := calculateCRC16(crcString) // คำนวณ CRC16
	qrStringWithCRC := qrString + crc // เพิ่ม CRC16 เข้าไปใน QR String

	// สร้าง QR Code เป็นภาพ PNG
	qrCodeFile := fmt.Sprintf("qr-%d.png", request.Amount)
	err := qrcode.WriteFile(qrStringWithCRC, qrcode.Medium, 256, qrCodeFile)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate QR code"})
		return
	}

	// ส่งข้อมูล QR String กลับไป
	c.JSON(http.StatusOK, gin.H{"qrString": qrStringWithCRC})
}

func generatePromptPayQR(promptPayID string, amount float64) string {
	// รูปแบบข้อมูลสำหรับ PromptPay QR
	payloadFormat := "000201"
	pointOfInitiation := "010212"
	merchantAccountInfo := fmt.Sprintf("29370016A000000677010111%s", promptPayID)
	transactionAmount := fmt.Sprintf("54%02d%.2f", len(fmt.Sprintf("%.2f", amount)), amount)
	countryCode := "5802TH"
	crc16 := "6304" // ส่วน CRC ยังไม่ถูกคำนวณ

	// รูปแบบ QR String ก่อนคำนวณ CRC16
	fullPromptPay := fmt.Sprintf("%s%s%s%s%s%s", 
		payloadFormat, pointOfInitiation, merchantAccountInfo, countryCode, transactionAmount, crc16)

	return fullPromptPay
}

func calculateCRC16(input string) string {
	const polynomial uint32 = 0x11021 // ใช้ uint32 แทน uint16 เพื่อรองรับค่าใหญ่
	var crc uint16 = 0xFFFF

	for _, b := range input {
		crc ^= uint16(b) << 8
		for i := 0; i < 8; i++ {
			if crc&0x8000 != 0 {
				crc = uint16((uint32(crc) << 1) ^ polynomial) // ใช้ uint32 ช่วยคำนวณ
			} else {
				crc <<= 1
			}
		}
	}
	crc &= 0xFFFF // ทำให้แน่ใจว่า CRC16 อยู่ในขอบเขต uint16
	return fmt.Sprintf("%04X", crc)
}
