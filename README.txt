การ Auto Run 
ใช้ libary : air
การติดตั้ง
1.เพิ่ม path เข้าไปใน Window
>> environment varibles
>> Edit the system environment varibles
>> Environment Varibles
>> ช่อง (User varibles for HP) -> Path -> Edit
>> New -> (C:\Users\HP\go\bin)
>> OK
>> ปิด cmd , VS code

2.go mod init github.com/xxxx/(folderxxx) //ทำทุกครั้ง
3.go install github.com/air-verse/air@latest
4.air

-------------------------------------------------------
การสร้าง Web Server ด้วย
fiber : go get github.com/gofiber/fiber/v2


-------------------------------------------------------
วิธีการเเก้ air ไม่รัน
>> กด Window -> Powershell
>> คลิ๊กขวา -> run as Administartor
>> Set-ExecutionPolicy -Scope CurrentUser -ExecutionPolicy RemoteSigned
 