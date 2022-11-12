Certificates service
===

Запуск сервиса:
"cmd/rpc/main.go"

Перед запуском необходимо:
Задать переменные окружения - "TEMPLATES_DIR", "CERTIFICATES_DIR",
    с путями хранения файлов шаблонов и сертификатов.
В системе должна быть установлена "wkhtmltopdf", программа конвертации HTML файлов в PDF.
    "wkhtmltopdf": https://wkhtmltopdf.org/downloads.html
    Исполняемый файл "wkhtmltopdf" должен находиться в каталоге запуска сервиса,
    или быть доступным по путям из переменной окружения "PATH"
Сгенерировать из ".proto" вспомогательные файлы gRPC: "protoc @protoc_options_generate.txt"

Требования к шаблонам:
В шаблоне могут содержаться следующие теги замены:
    {{.CourseName}}
    {{.CourseType}}
    {{.CourseHours}}
    {{.CourseDate}}
    {{.CourseMentors}}
    {{.StudentFirstname}}
    {{.StudentLastname}}
Шаблоны с любыми другими тегами замены будут отклонены валидатором.
