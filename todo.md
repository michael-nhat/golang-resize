##### Todo:
    - [v] init
    - [v] reload
    - [v] return image
    - [x] process params
    - [x] resize

go simple resize

Requests per second:    305.76 [#/sec] (mean)
Time per request:       3270.554 [ms] (mean)
Time per request:       3.271 [ms] (mean, across all concurrent requests)
Transfer rate:          7424.20 [Kbytes/sec] received

fill resize local
Requests per second:    342.81 [#/sec] (mean)
Time per request:       2917.047 [ms] (mean)
Time per request:       2.917 [ms] (mean, across all concurrent requests)
Transfer rate:          3963.43 [Kbytes/sec] received

fill resize remote test
Requests per second:    92.67 [#/sec] (mean)
Time per request:       10790.997 [ms] (mean)
Time per request:       10.791 [ms] (mean, across all concurrent requests)
Transfer rate:          1054.84 [Kbytes/sec] received

resize wh local
Requests per second:    1451.57 [#/sec] (mean)
Time per request:       688.909 [ms] (mean)
Time per request:       0.689 [ms] (mean, across all concurrent requests)
Transfer rate:          32417.95 [Kbytes/sec] received

resize wh remote
Requests per second:    232.29 [#/sec] (mean)
Time per request:       4304.970 [ms] (mean)
Time per request:       4.305 [ms] (mean, across all concurrent requests)
Transfer rate:          5148.71 [Kbytes/sec] received


y tuong dich vu image:
cache image resize tren hdd
su dung cronjob de compare va date khi can thien
va cung xoa voi file khong truy cap nhieu, chi de lai file truy cap nhieu, lan dau tien anh se hoi cham
cho nen can xem truoc de no resize
check file info and compare file la bat buoc??

luon luon phai có 2 lần check fileinfo? file gốc và target?

neu là file tinh trong bai post khong bi chinh sua thi co the tra ve luon

có the là nó lay anh dau tien trong post 
nhung anh sau se la crop cua link anh do??

nhu vay phai gen ra file la nhanh nhat, nhung khong tiet kiem bo nho

co the co mot service thong ke truy cap, tu dong xoa nhung file lau khong truy cap  (2 tuan chang han) va nhung file co tan suat thap ( chac la phai co cong thuc, quy tac , thuat toan cu the)

tactic: cold cache, hot cache, cron job, => cpu (also for waste, cpu free also waste)


