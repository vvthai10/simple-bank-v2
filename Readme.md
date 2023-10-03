# Demo backend GO | simple-bank

### vvtahi


**Mô tả:** Xây dựng một hệ thống theo "Clean Architecture" bằng ngôn ngữ Golang. Các yêu cầu trong project:
- Sử dụng package GORM + Postgres DB
- Xây dựng logic cho cơ chế transaction
- Hiểu và giải thích được các nguyên nhân một vài trường hợp deadlock xảy ra

**Kết quả:**
- Tạo các API tương tác với cơ sở dữ liệu
- Tạo các test=case thể hiện các trường hợp đặc biệt của transaction
- Giải thích được một vài trường hợp deadlock

**Tổng quan về các phần của project:** 
- User: Thể hiện các thuộc tính của người dùng, được cung cấp các cơ chế create, get, update và delete
- Account: Mỗi User sẽ có 1 hoặc nhiều account(**chưa tạo logic xử lý về quy định các yêu cầu về giới hạn account**), cung cấp các cơ chế create và update
- Entry: Lưu thông tin về các lần cập nhật số dư của các account
- Transfer: Lưu thông tin về các lần giao dịch qua lại giữa các account - **chứa các xử lý về transaction**

**Giải thích, phân tích về phần transaction**
- Yêu cầu: Một tài khoản có thể thực hiện nhiều giao dịch đồng thời với nhau, khi đó sẽ phải có sự cập nhật qua lại giữa các bảng **Entry** và **Account**. Khi đó cần phải tạo ra **transaction** để đảm bảo các quá trình trên sẽ được thực hiện 1 cánh đúng đắn. Với trường hợp này sẽ có 5 bước khi thực hiện chuyển tiền từ tài khoản A sang tài khoản B
<ul>
    <ul>
        <li>Create 1 dòng trong bảng transaction từ A, sang B với M đồng</li>
        <li>Create 1 dòng trong Entry ghi lại tài khoản A bị trừ M đồng</li>
        <li>Create 1 dòng trong Entry ghi lại tài khoản B nhận được M đồng</li>
        <li>Update account A bị trừ M đồng</li>
        <li>Update account B được cộng M đồng</li>
    </ul>
</ul>

- Transaction đảm bảo về 4 tiêu chí:
    <ul>
        <li>Atomicity - tính nguyên tử</li>
        <li>Consistency - tính nhất quán</li>
        <li>Isolation - tính cô đọng</li>
        <li>Durability - tính bền vững</li>
    </ul>
- Một transaction khi thực hiện sẽ có 2 trường hợp: BEGIN --- COMMIT và BEGIN --- ROLLBACK

**Các trường hợp lỗi / deadlock với transaction**
- Với việc thực hiện transfer trên thì có các bước tạo các dòng mới thì sẽ không gặp các trường hợp deadlock hoặc xử lý sai. Các xử lý sẽ gặp lỗi khi đến việc chọn và cập nhật các dòng trong bảng account, một vài tính huống mà gặp trong project này
    <ul>
        <li>- Với cách lấy và cập nhật account bình thường, thì sẽ gặp tình trạng không chặn các truy cập tại dòng đó trong nhiều transaction cùng lúc. Khi đó dẫn tới việc nhiều transaction cùng lấy giá trị dòng đó tại 1 thời điểm và dẫn đến cập nhật không đúng --> <b>Cách xử lý SELECT FOR UPDATE</b></li>
        <li>Trường hợp block khi SELECT FOR UPDATE vào bảng account bà bị chặn bởi lệnh INSERT vào bảng transfer. <b>Tại sao SELECT lại đợi INSERT? Vì đang có mối liên hệ khóa ngoại giữa 2 bảng, việc block đảm bảo cho sự đồng nhất dữ liệu</b> ---> <b>Cách xử lý là sử dụng hàm SELECT NO KEY UPDATE để đảm bảo là key ủa các bảng sẽ không bị thay đổi</b></li>
        <li>Trường hợp block khi thực hiện đồng thời 2 giao dịnh từ A sang B và từ B sang A: Khi này sẽ là trường hợp transaction 1 đợi transaction 2 và transaction 2  cũng đợi transaction 1 --> <b>Hướng giải quyết, làm cho các transaction sẽ đi từ 1 chiều từ A sang B nhưng vẫn đảm bảo sự đúng đắn trong việc chuyển tiền</b></li>
    </ul>
