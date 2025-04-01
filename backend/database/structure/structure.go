package structure

type User struct {
	UserId           int    `json:"user_id"`
	Username         string `json:"username"`
	Email            string `json:"email"`
	Password         string `json:"password"`
	FirstName        string `json:"first_name"`
	LastName         string `json:"last_name"`
	Gender           string `json:"gender"`
	DateOfBirth      string `json:"date_of_birth"`
	Age              int    `json:"age"`
	RegistrationDate string `json:"registration_date"`
	ProfilePicture   string `json:"profile_picture"`
	IsConnected      int    `json:"is_connected"`
	Role             int    `json:"role"`
}

type Post struct {
	PostId          int    `json:"post_id"`
	UserId          int    `json:"user_id"`
	Title           string `json:"title"`
	PublicationDate string `json:"publication_date"`
	Categorie       string `json:"categorie"`
	Message         string `json:"message"`
	NbReply         int    `json:"nb_reply"`
	NbLike          int    `json:"nb_like"`
	NbDislike       int    `json:"nb_dislike"`
	IsLiked         bool   `json:"is_liked"`
}

type Comment struct {
	CommentId       int    `json:"comment_id"`
	UserId          int    `json:"user_id"`
	PostId          int    `json:"post_id"`
	Title           string `json:"title"`
	PublicationDate string `json:"publication_date"`
	Message         string `json:"message"`
}

type Categorie struct {
	CategorieId int    `json:"categorie_id"`
	Name        string `json:"name"`
}

type Message struct {
	MessageId  int    `json:"message_id"`
	Sender     string `json:"sender"`
	Recipient  string `json:"recipient"`
	Message    string `json:"message"`
	Timestamp  string `json:"timestamp"`
	ReadStatus bool   `json:"read_status"`
}

type Notifications struct {
	Type         string
	Notification map[string]int
	Count        int
	Recipient    string
}

type UserList struct {
	Type      string
	UserList  []User
	Sender    string
	Recipient string
}
