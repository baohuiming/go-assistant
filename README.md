Guide:
1. [Get A OpenWeatherMap Api Key](https://home.openweather.co.uk/api_keys).
2. Get LLM api url and key.
3. Create a `config.yaml` file.

Example of `config.yaml`:
```yaml
weather:
    api_key: your_openweathermap_api_key
    latitude: your_latitude
    longitude: your_longitude

dlna:
    device_name: "DLNA-Feixun-R1"

llm:
    base_url: "https://api.siliconflow.cn/v1"
    api_key: "sk-********"
    model: "Qwen/Qwen2-7B-Instruct"
    prompt: >
        你是一个AI智能家庭助理，这是自动获取得到的未来18个小时的天气情况，如果有下雨/下雪/恶劣天气的话就提醒主人，*回答尽量简短、活泼、一句话*：
        """
        %s
        """
```
