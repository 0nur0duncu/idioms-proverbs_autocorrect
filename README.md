# Otomatik Kelime Tamamlama Servisi

Bu proje, Excel dosyasından okunan kelimeler için otomatik tamamlama önerileri sunan bir web servisidir. Trie veri yapısı kullanılarak hızlı arama ve öneriler sağlanır.

## Özellikler

- Excel dosyasından kelime okuma
- Trie veri yapısı ile etkili arama
- REST API desteği
- Web arayüzü
- Cross-Origin Resource Sharing (CORS) desteği

## Teknolojiler

- Go
- Fiber web framework
- Excelize (Excel okuma)
- HTML/CSS/JavaScript

## Kurulum

1. Gerekli Go paketlerini yükleyin:
```bash
go get github.com/gofiber/fiber/v2
go get github.com/xuri/excelize/v2
```

2. Projeyi çalıştırın:
```bash
go run main.go
```

3. Web tarayıcısında açın:
```
http://localhost:8080
```

## API Kullanımı

### Kelime Önerileri Alma

```
GET /autocomplete?prefix={aranacak_kelime}
```

Örnek yanıt:
```json
{
    "prefix": "ara",
    "suggestions": ["arada", "aralık", "araya"]
}
```

## Excel Dosya Yapısı

Kelimeler Excel dosyasının ilk sütununda bulunmalıdır. İlk satır başlık olarak kabul edilir ve atlanır.

## Lisans

Bu proje MIT lisansı ile lisanslanmıştır. Daha fazla bilgi için [LICENSE](LICENSE) dosyasını inceleyebilirsiniz.
