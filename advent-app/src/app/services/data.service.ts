import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { catchError } from 'rxjs/operators';
import { throwError } from 'rxjs';

export interface Data {
  answer1: string|number;
  answer2: string|number;
  metaData: any;
}

export interface AdventResponse {
  status: number;
  data: Data;
}

export interface AdventResponseError {
  status: number;
  error: string;
}

@Injectable({
  providedIn: 'root'
})
export class DataService {

  constructor(
    private http: HttpClient
  ) { }

  getDayAnswer(day: number) {
    return this.http.get<AdventResponse>(`http://localhost:8084/advent/day/${day}`)
      .pipe(
        catchError(error => {
          console.error(error);
          return throwError(() => error.error);
        })
      );
  }
}
