import { Component, Input } from '@angular/core';

@Component({
  selector: 'answer',
  templateUrl: './answer.component.html',
  styleUrls: ['./answer.component.scss']
})

export class AnswerComponent {

  @Input() answer: string | number | undefined = undefined
  @Input() label: string = "";

  constructor() { }
}
