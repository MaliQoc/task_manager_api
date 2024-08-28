## API Uç Noktaları (Örneğin Postman ile test ederken)

- **Ana Sayfa:** `GET /` - Ana sayfa mesajı döner.
- **Görevleri Listele:** `GET /tasks` - Tüm görevlerin listesini döner.
- **Görev Detayı:** `GET /tasks/:id` - Belirli bir görev ID'sine göre görev detayını döner.
- **Yeni Görev Ekle:** `POST /tasks` - Yeni bir görev oluşturur.

## Kurulum (Visual Studio Code terminalinde)

1. Bağımlılıkları yükleyin:
   ```terminal
   go mod tidy
   
2. Uygulamayı başlatın:
   ```terminal
   go run main.go

Uygulama localhost:8080 üzerinde çalışacaktır. Örneğin http://localhost:8080/tasks yazıp Postman'da istediğiniz işleme göre (GET, POST) eylemlerinizi gerçekleştirebilirsiniz.
