package internal

/* header user-agent 防护，并且采用散列
// 定义阈值和时间窗口
const threshold = 10      // 请求频率阈值
const windowSizeInSeconds = 60  // 时间窗口大小（以秒为单位）

// 使用map来存储每个散列后的User-Agent的请求次数
var requestHistory = make(map[string]int)

func generateKey(userAgent string) string {
	hasher := sha1.New()
	hasher.Write([]byte(userAgent))
	return hex.EncodeToString(hasher.Sum(nil))
}

func checkCCAttack(w http.ResponseWriter, r *http.Request) {
	userAgent := r.Header.Get("User-Agent")

	if userAgent != "" {
		key := generateKey(userAgent)
		requestHistory[key]++

		// 如果请求次数超过阈值，则返回拒绝访问
		if requestHistory[key] > threshold {
			http.Error(w, "Too Many Requests", http.StatusTooManyRequests)
			return
		}
	}

	w.Write([]byte("OK"))
}
}

func main() {
	http.HandleFunc("/", checkCCAttack)
	http.ListenAndServe(":8080", nil)
}
*/

/* 确保唯一
func generateKey(userAgent string) string {
	uniqueID := time.Now().UnixNano() + rand.Int63n(1000)  // 结合时间戳和随机数生成唯一ID
	key := strings.Join([]string{userAgent, strconv.FormatInt(uniqueID, 10)}, "|")
	hasher := sha1.New()
	hasher.Write([]byte(key))
	return hex.EncodeToString(hasher.Sum(nil))
}
*/

/* cookie 防护
# 定义阈值和时间窗口
threshold = 10  # 请求频率阈值
window_size = 60  # 时间窗口大小（以秒为单位）

# 使用字典默认值为列表来存储每个会话标识的请求时间戳
request_history = defaultdict(list)

@app.route('/')
def index():
    client_id = request.cookies.get('client_id')  # 获取客户端的会话标识（Cookie名称为'client_id'）

    if client_id:
        current_time = time.time()

        # 获取当前客户端的请求时间戳列表
        timestamps = request_history[client_id]

        # 清理超过时间窗口的旧时间戳
        while timestamps and timestamps[0] <= current_time - window_size:
            timestamps.pop(0)

        # 添加当前时间戳到列表
        timestamps.append(current_time)

        # 如果请求数超过阈值，则返回拒绝访问
        if len(timestamps) > threshold:
            return "Too many requests. Please try again later."

    # 设置新的Cookie作为会话标识
    response = make_response("Hello, World!")
    response.set_cookie('client_id', generate_client_id())  # 生成唯一的会话标识并设置为Cookie

    return response

def generate_client_id():
    # 在实际应用中，你可以根据需要生成自己的会话标识，例如使用UUID等
    return str(uuid.uuid4())
*/
