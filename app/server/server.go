package server

import (
	"context"

	valid "github.com/go-ozzo/ozzo-validation/v4"
	validIs "github.com/go-ozzo/ozzo-validation/v4/is"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"gus_certificates/app/certgenerator"
	certSPb "gus_certificates/protobuf/transport/certificate"
	"gus_certificates/utils/pdfgenerator"
	"gus_certificates/utils/qrgenerator"
	"gus_certificates/utils/storage"
)

var fileNameRule = []valid.Rule{
	valid.Required,
	valid.RuneLength(5, 255),
	validIs.ASCII,
}

const certificateFileExtension = ".pdf"

type certificateServer struct {
	certSPb.UnimplementedCertificateServer

	// Генерация сертификата в HTML, генерация ID сертификата.
	certGen *certgenerator.CertGenerator

	// Конвертация сертификата в формат PDF.
	pdfGen pdfgenerator.PdfGenerator

	// Работа с локальным файловым хранилищем.
	strgLoc storage.Storage

	// Генерация QR кодов.
	qrGen *qrgenerator.QrGenerator
}

func NewCertificateServer() (*certificateServer, error) {
	pdfGen, err := pdfgenerator.New()
	if err != nil {
		return nil, err
	}

	strgLoc, err := storage.NewLocal()
	if err != nil {
		return nil, err
	}

	server := &certificateServer{}
	server.certGen = &certgenerator.CertGenerator{}
	server.pdfGen = pdfGen
	server.strgLoc = strgLoc
	server.qrGen = &qrgenerator.QrGenerator{}

	return server, nil
}

func (c *certificateServer) fillData(course *certSPb.CourseMessage, student *certSPb.StudentMessage) {
	c.certGen.SetCourseName(course.GetCourseName())
	c.certGen.SetCourseType(course.GetCourseType())
	c.certGen.SetCourseHours(course.GetHours())
	c.certGen.SetCourseDate(course.GetDate())
	c.certGen.SetCourseMentors(course.GetMentors())
	c.certGen.SetStudentFirstname(student.GetFirstname())
	c.certGen.SetStudentLastname(student.GetLastname())
}

func (c *certificateServer) IssueCertificate(ctx context.Context, req *certSPb.IssueCertificateReq) (*certSPb.IssueCertificateResp, error) {
	// Валидация имени шаблона.
	templateName := req.GetTemplateName()
	err := valid.Validate(templateName, fileNameRule...)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "%s: %q %v", "IssueCertificate", templateName, err)
	}

	// Заполнение и валидация данных о курсе и студенте.
	c.fillData(req.GetCourse(), req.GetStudent())
	err = c.certGen.ValidateData()
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "%s: %v", "IssueCertificate", err)
	}

	// Генерация ID сертификата.
	certificateId := c.certGen.GenerateID()

	// Генерация QR-Code на линк сертификата.
	qrCodeLinkPNG, err := c.qrGen.GenerateQrPNG(certificateId) // Пока не реализовано получение линка передается имя сертификата
	if err != nil {
		return nil, status.Errorf(codes.FailedPrecondition, "%s: %v", "IssueCertificate", err)
	}
	c.certGen.SetQrCodeLink(qrCodeLinkPNG)

	// Получение шабона из хранилища.
	template, err := c.strgLoc.GetTemplate(templateName)
	if err != nil {
		return nil, status.Errorf(codes.FailedPrecondition, "%s: %q %v", "IssueCertificate", templateName, err)
	}

	// Генерация сертификата в формате HTML.
	certificate, err := c.certGen.GenerateCertHTML(template)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "%s: %v", "IssueCertificate", err)
	}

	// Конвертация сертификата в формат PDF.
	certificatePDF, err := c.pdfGen.RenderHtmlToPdf(certificate)
	if err != nil {
		return nil, status.Errorf(codes.FailedPrecondition, "%s: %v", "IssueCertificate", err)
	}

	// Сохранение сертификата в хранилище.
	err = c.strgLoc.SaveCertificate(certificateId+certificateFileExtension, certificatePDF)
	if err != nil {
		return nil, status.Errorf(codes.FailedPrecondition, "%s: %v", "IssueCertificate", err)
	}

	resp := &certSPb.IssueCertificateResp{Id: certificateId}
	return resp, nil
}

func (c *certificateServer) GetCertificateFileByID(ctx context.Context, req *certSPb.GetCertificateFileByIDReq) (*certSPb.GetCertificateFileByIDResp, error) {
	// Валидация имени сертификата.
	certificateId := req.GetId()
	err := valid.Validate(certificateId, fileNameRule...)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "%s: %q %v", "GetCertificateFileByID", certificateId, err)
	}

	// Получение сертификата из хранилища.
	certificate, err := c.strgLoc.GetCertificate(certificateId + certificateFileExtension)
	if err != nil {
		return nil, status.Errorf(codes.FailedPrecondition, "%s: %q %v", "GetCertificateFileByID", certificateId, err)
	}

	resp := &certSPb.GetCertificateFileByIDResp{Certificate: certificate}
	return resp, nil
}

func (c *certificateServer) GetCertificateLinkByID(ctx context.Context, req *certSPb.GetCertificateLinkByIDReq) (*certSPb.GetCertificateLinkByIDResp, error) {
	// Валидация имени сертификата.
	certificateId := req.GetId()
	err := valid.Validate(certificateId, fileNameRule...)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "%s: %q %v", "GetCertificateLinkByID", certificateId, err)
	}

	// Получение линка на сертификат.
	certificateLink := certificateId // Пока не реализовано получение линка передается имя сертификата
	// certificateFullPath, err := c.strgLoc.GetCertificatePath(certificateId + certificateFileExtension)
	// if err != nil {
	// 	return nil, status.Errorf(codes.FailedPrecondition, "%s: %q, %v", "GetCertificateLinkByID", certificateId, err)
	// }

	resp := &certSPb.GetCertificateLinkByIDResp{Link: certificateLink}
	// resp := &certSPb.GetCertificateLinkByIDResp{Link: certificateFullPath}
	return resp, nil
}

func (c *certificateServer) AddTemplate(ctx context.Context, req *certSPb.AddTemplateReq) (*certSPb.AddTemplateResp, error) {
	// Валидация имени шаблона.
	templateName := req.GetTemplateName()
	err := valid.Validate(templateName, fileNameRule...)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "%s: %q %v", "AddTemplate", templateName, err)
	}

	// Проверка корректности файла шаблона.
	templateByte := req.GetTemplate()
	err = c.certGen.CheckTemplateHTML(templateByte)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "%s: %q %v", "AddTemplate", templateName, err)
	}

	// Сохранение шаблона в хранилище.
	err = c.strgLoc.SaveTemplate(templateName, templateByte)
	if err != nil {
		return nil, status.Errorf(codes.FailedPrecondition, "%s: %q %v", "AddTemplate", templateName, err)
	}

	resp := &certSPb.AddTemplateResp{Status: &certSPb.Status{Code: int32(codes.OK)}}
	return resp, nil
}

func (c *certificateServer) DeleteTemplate(ctx context.Context, req *certSPb.DeleteTemplateReq) (*certSPb.DeleteTemplateResp, error) {
	// Валидация имени шаблона.
	templateName := req.GetTemplateName()
	err := valid.Validate(templateName, fileNameRule...)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "%s: %q %v", "DelTemplate", templateName, err)
	}

	// Удаление шаблона из хранилища.
	err = c.strgLoc.DeleteTemplate(templateName)
	if err != nil {
		return nil, status.Errorf(codes.FailedPrecondition, "%s: %q %v", "DelTemplate", templateName, err)
	}

	resp := &certSPb.DeleteTemplateResp{Status: &certSPb.Status{Code: int32(codes.OK)}}
	return resp, nil
}
