package server

import (
	"context"
	"gus_certificates/app/certgenerator"
	certSPb "gus_certificates/protobuf/transport/certificate"
	"gus_certificates/utils/pdfgenerator"
	"gus_certificates/utils/storage"

	valid "github.com/go-ozzo/ozzo-validation/v4"
	validIs "github.com/go-ozzo/ozzo-validation/v4/is"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var fileNameRule = []valid.Rule{
	valid.Required,
	valid.RuneLength(5, 255),
	validIs.ASCII,
}

type certificateServer struct {
	certSPb.UnimplementedCertificateServer

	// Генерация сертификата в HTML, генерация ID сертификата.
	certGen *certgenerator.CertGenerator

	// Конвертация сертификата в формат PDF.
	pdfGen pdfgenerator.PdfGenerator

	// Работа с локальным файловым хранилищем.
	strgLoc storage.Storage
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

	return server, nil
}

func (c *certificateServer) IssueCertificate(ctx context.Context, req *certSPb.IssueCertificateReq) (*certSPb.IssueCertificateResp, error) {
	// Валидация имени шаблона.
	templateName := req.GetTemplateName()
	err := valid.Validate(templateName, fileNameRule...)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "%s: %q %v", "IssueCertificate", templateName, err)
	}

	// Заполнение и валидация данных о курсе и студенте.
	course := req.GetCourse()
	student := req.GetStudent()
	c.certGen.SetCourseName(course.GetCourseName())
	c.certGen.SetCourseType(course.GetCourseType())
	c.certGen.SetCourseHours(course.GetHours())
	c.certGen.SetCourseDate(course.GetDate())
	c.certGen.SetCourseMentors(course.GetMentors())
	c.certGen.SetStudentFirstname(student.GetFirstname())
	c.certGen.SetStudentLastname(student.GetLastname())
	err = c.certGen.ValidateData()
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "%s: %v", "IssueCertificate", err)
	}

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

	// Генерация ID/имени сертификата (ID + ".pdf").
	nameCertificateIdPDF := c.certGen.GenerateID() + ".pdf"

	// Сохранение сертификата в хранилище.
	err = c.strgLoc.SaveCertificate(nameCertificateIdPDF, certificatePDF)
	if err != nil {
		return nil, status.Errorf(codes.FailedPrecondition, "%s: %v", "IssueCertificate", err)
	}

	resp := &certSPb.IssueCertificateResp{Id: nameCertificateIdPDF}
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
	certificate, err := c.strgLoc.GetCertificate(certificateId)
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

	// Получение полного пути до файла сертификата в хранилище.
	certificateFullPath, err := c.strgLoc.GetCertificatePath(certificateId)
	if err != nil {
		return nil, status.Errorf(codes.FailedPrecondition, "%s: %q, %v", "GetCertificateLinkByID", certificateId, err)
	}

	resp := &certSPb.GetCertificateLinkByIDResp{Link: certificateFullPath}
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

func (c *certificateServer) DelTemplate(ctx context.Context, req *certSPb.DelTemplateReq) (*certSPb.DelTemplateResp, error) {
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

	resp := &certSPb.DelTemplateResp{Status: &certSPb.Status{Code: int32(codes.OK)}}
	return resp, nil
}
