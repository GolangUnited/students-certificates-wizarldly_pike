# Certificates service
===

### Запуск сервиса в Docker:
#### Создание образа и генерация **protobuf** файлов из **".proto"**:
`docker image build -f Dockerfile.proto -t cert-proto .`

#### Сохранение в локальном проекте **protobuf** файлов из Docker контейнера **cert-proto**:
`docker container run --rm -v $PWD/protobuf/transport/certificate:/result cert-proto` // Linux  
`docker container run --rm -v %CD%/protobuf/transport/certificate:/result cert-proto` // Windows

#### Создание образа для запуска проекта:
`docker image build -f Dockerfile.app -t cert-app .`

#### Запуск проекта с сохранением в Local Storage:
`docker container run --rm --env-file env.list -p 1234:1234 -it -v certificate-storage:/storage cert-app`   
В файле **env.list** задаются переменные откружения необходимые для запуска сервиса.  
Тип Storage задается перемнной **STORAGE=local**, переменную **STORAGE** можно опустить, значение **local** будет присвоено по умолчанию.  
Монтируется Docker VOLUME: **certificate-storage** на локальной машине/сервере, для сохранения шаблонов и сертификатов.

#### Запуск проекта с сохранением в AWS S3 :
`docker container run --rm --env-file env.list --env-file s3.env -p 1234:1234 -it cert-app`  
В файле **env.list** обязательно задается переменная окружения  **STORAGE=s3**.  
В файле **s3.env** обязательно задаються переменные окружения:  
```
S3_ENDPOINT=... // s3.amazonaws.com - если планируем использовать облачное хранилище от Amazon.
S3_BUCKET_NAME=... // Имя существующего Bucket, на который у вас есть права записи.
ACCESS_KEY_ID=...
SECRET_ACCESS_KEY=...
```
### Требования к шаблонам:
В шаблоне могут содержаться следующие теги замены:
```
{{.CourseName}}
{{.CourseType}}
{{.CourseHours}}
{{.CourseDate}}
{{.CourseMentors}}
{{.StudentFirstname}}
{{.StudentLastname}}
{{.QrCodeLink}}
```
Шаблоны с любыми другими тегами замены будут отклонены валидатором.  
Вместо тега замены `{{.QrCodeLink}}` будет вставлен **QR** код в формате **PNG**: ссылка на сертификат.

### Пример простого HTML шаблона:
```
<html><body><h1 style="color:red;">Test html color<h1>
<p>{{.CourseName}}</p>
<p>{{.CourseType}}</p>
<p>{{.CourseHours}}</p>
<p>{{.CourseDate}}</p>
<p>{{.CourseMentors}}</p>
<p>{{.StudentFirstname}}</p>
<p>{{.StudentLastname}}</p>
<p><img src={{.QrCodeLink}} width="128" height="128"></p>
</body></html>
```
