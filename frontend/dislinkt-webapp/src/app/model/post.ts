export interface Post {
    id        : string,
	username  : string,
	text      : string,
	liked     : string[],
	disliked  : string[],
	CreatedOn : string,
	imageUrl  : string,
	comments  : Comment[]
}

export class Comment {
    username! : string;
    text !    : string;
}
