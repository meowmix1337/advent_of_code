import { Component } from '@angular/core';
interface Day {
  dayNumber: number;
  title: string;
  description: string;
}

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.scss']
})
export class AppComponent {
  title = 'advent-app';

  days: Day[] = [{
    dayNumber: 1,
    title: 'Calories Count',
    description: 'Sum the total calories'
  }, {
    dayNumber: 2,
    title: 'Rock Paper Scissors',
    description: 'Use proper strategy to win!'
  }, {
    dayNumber: 3,
    title: 'Rucksack Reorganization',
    description: 'Sum the priorities'
  }, {
    dayNumber: 4,
    title: 'Camp Cleanup',
    description: 'What overlaps!!'
  }, {
    dayNumber: 5,
    title: 'Supply Stacks',
    description: 'What ends up on the top of the stacks!!'
  }, {
    dayNumber: 6,
    title: 'Tuning Trouble',
    description: 'Find the start of the packet marker!!'
  }, {
    dayNumber: 7,
    title: 'No Space Left On Device',
    description: 'Find the correct directory to delete!'
  }, {
    dayNumber: 8,
    title: 'Treetop Tree House',
    description: 'Find the best location for the tree house!'
  }, {
    dayNumber: 9,
    title: 'Rope Bridge',
    description: 'Snake'
  }, {
    dayNumber: 10,
    title: 'Cathode-Ray Tube',
    description: 'CRT and CPU cycles!'
  }, {
    dayNumber: 11,
    title: 'Monkey in the Middle',
    description: 'What is the monkey business level?!'
  }, {
    dayNumber: 12,
    title: 'Hill Climbing Algorithm',
    description: 'Fewest steps!'
  }];

  days_2023: Day[] = [{
    dayNumber: 1,
    title: 'Trebuchet?!',
    description: 'Find the calibration Values!'
  }, {
    dayNumber: 2,
    title: 'Cube Conundrum',
    description: 'Sum of IDs!'
  }, {
    dayNumber: 3,
    title: 'Gear Ratios',
    description: 'Find the Gear Ratios!'
  }, {
    dayNumber: 4,
    title: 'Scratchcards',
    description: 'Sum of your total points!'
  }, {
    dayNumber: 5,
    title: 'If You Give A Seed A Fertilizer',
    description: 'This might take some time'
  }, {
    dayNumber: 6,
    title: 'Wait For It',
    description: 'Race time!'
  }, {
    dayNumber: 7,
    title: 'Camel Cards',
    description: 'What are your winnings!'
  }, {
    dayNumber: 8,
    title: 'Haunted Wasteland',
    description: 'Least Common Multiple!'
  }];
}
