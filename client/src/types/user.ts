export interface IAPIResponse {
  message: string;
  user?: IUser;
  redirect?: boolean;
}

export interface IUser {
  ID: number;
  register_no: string;
  email: string;
  role: string;
  isVerified: boolean;
  fullname: string;
  CreatedAt: Date;
}
