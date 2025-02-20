export interface IClassesAPIResponse {
  classes: IClasses[];
  message: string;
  success: boolean;
}

export interface IClasses {
  ID: number;
  userId: number;
  name: string;
  maxClasses: number;
  totalTaken: number;
  attendedClasses: number;
  start_time: string;
  end_time: string;
  status: boolean;
}

// subjects summary
export interface ISummaryAPIResponse {
  message: string;
  success: boolean;
  summary: Summary;
}

export interface Summary {
  overall_percentage: number;
  subjects: Subjects;
  total_attended: number;
  total_classes: number;
  total_missed: number;
}

export interface Subjects {
  ADA: SubjectSummary;
  "ADA Lab": SubjectSummary;
  ENG: SubjectSummary;
  IC: SubjectSummary;
  IT: SubjectSummary;
  "IT Lab": SubjectSummary;
  LANG: SubjectSummary;
  OE: SubjectSummary;
  SE: SubjectSummary;
}

export interface SubjectSummary {
  attended_classes: number;
  max_classes: number;
  percentage: number;
  remaining: number;
}
