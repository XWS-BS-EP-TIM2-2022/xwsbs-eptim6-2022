export interface Post {
    ID        : string,
	Username  : string,
	Text      : string,
	Likes     : number,
	Liked     : string[],
	Dislikes  : number,
	Disliked  : string[],
	CreatedOn : string,
	ImageUrl  : string,
	Comments  : Comment[]
}

export interface Comment {
    Username : string,
    Text     : string
}