interface IAPIResponse {
  messaage: string;
  user?: IUser;
}

interface IUser {
  ID: number;
  Username: string;
  Email: string;
  Role: string;
  IsVerified: boolean;
  Fullname: string;
  CreatedAt: Date;
}
