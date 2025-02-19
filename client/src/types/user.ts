export interface IAPIResponse {
  message: string;
  user?: IUser;
  redirect?: boolean;
}

export interface IUser {
  ID: number;
  Username: string;
  Email: string;
  Role: string;
  IsVerified: boolean;
  Fullname: string;
  CreatedAt: Date;
}
