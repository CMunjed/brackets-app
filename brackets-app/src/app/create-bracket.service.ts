import { Injectable } from '@angular/core';
import { BehaviorSubject } from 'rxjs';
/*
@Injectable({
  providedIn: 'root'
})
*/
class Match {
  Member1:     string;
	Member2:     string;
	Member1Wins: number;
	Member2Wins: number;

  constructor(m1: string, m2: string, m1w: number, m2w: number){
    this.Member1 = m1;
    this.Member2 = m2;
    this.Member1Wins = m1w;
    this.Member2Wins = m2w;
  }
}

class Bracket {
  Teams: number
  TeamsList: string[]
  MatchList: Match[]
  Rounds: number
  isDouble: boolean

  constructor(r: number){
    this.Teams = r;
    this.isDouble = false;
    const m: Match[] = [];
    const t: string[] = [];
    this.MatchList = m;
    this.TeamsList = t;
    this.Rounds = Math.ceil(Math.log2(r)) + 1;
  }
}

export class CreateBracketService {
  winner:string = "";
  currentBracket = 0;
  bl: Bracket[] = [];
  b = new Bracket(0);
  
  getBracketList() {
    return this.bl;
  }
  getMatchList() {
    return this.b.MatchList;
  }
  getTeamsList() {
    return this.b.TeamsList;
  }
  getTeams() {
    return this.b.Teams.toString();
  }

  setTeamName(team:string, newName:string) {
    if(this.winner == team){
      this.winner = newName;
    }
    for(let i = 0; i < this.bl.length; i++){
      for(let j = 0; j < this.bl[i].MatchList.length; j++){
        if(this.bl[i].MatchList[j].Member1 == team){
          this.bl[i].MatchList[j].Member1 = newName;
          continue;
        }
        else if(this.bl[i].MatchList[j].Member2 == team){
          this.bl[i].MatchList[j].Member2 = newName;
          continue;
        }
      }
    }
  }

  getWinner(){
    return this.winner;
  }


  createBracket(Teams: number) {
    // simple check, title must be at least 1 char
    this.b.Teams = Teams;
    console.log(Teams)
    for (let i = 1; i <= Teams; i++) {
      this.b.TeamsList.push("Team " + i.toString());
    }
    for (let i = 0; i < Teams; i++) {
      console.log(this.b.TeamsList[i]);
    }

    const shuffledTeams = this.shuffle(this.b.TeamsList);
  
    for (let i = 0; i < shuffledTeams.length; i += 2) {
      const team1 = shuffledTeams[i];
      const team2 = shuffledTeams[i + 1];
      const newMatch = new Match(team1, team2, 0, 0);
      this.b.MatchList.push(newMatch);
    }


    this.bl.push(this.b);
    var t = this.b.Teams;
    this.b.Rounds = Math.ceil(Math.log2(this.b.Teams)) + 1;
    for (let i = 1; i <= this.b.Rounds; i++) {
      if(this.bl.length == this.b.Rounds - 1){
        break;
      }
      t = Math.floor(t/2);
      if(t == 1){
        var bTemp = new Bracket(1);
        bTemp.TeamsList.push("");

        const shuffledTeams = this.shuffle(bTemp.TeamsList);
  
        for (let i = 0; i < shuffledTeams.length; i += 2) {
          bTemp.MatchList.push(new Match("","",0,0));
        }

        this.bl.push(bTemp);
      }
      else if(t % 2 != 0){
        t = t + 1;
        var bTemp = new Bracket(t);
        for (let i = 1; i <= t; i++){
          bTemp.TeamsList.push("");
        }

        const shuffledTeams = this.shuffle(bTemp.TeamsList);
  
        for (let i = 0; i < shuffledTeams.length; i += 2) {
          bTemp.MatchList.push(new Match("","",0,0));
        }

        this.bl.push(bTemp);
      }
      else{
        var bTemp = new Bracket(t);
        for (let i = 1; i <= t; i++){
          bTemp.TeamsList.push("");
        }

        const shuffledTeams = this.shuffle(bTemp.TeamsList);
  
        for (let i = 0; i < shuffledTeams.length; i += 2) {
          bTemp.MatchList.push(new Match("","",0,0));
        }

        this.bl.push(bTemp);
      }
    }
  }

  createBracketDouble(Teams: number) {
    // simple check, title must be at least 1 char
    this.b.isDouble = true;
    this.b.Teams = Teams;
    console.log(Teams)
    for (let i = 1; i <= Teams; i++) {
      this.b.TeamsList.push("Team " + i.toString());
    }
    for (let i = 0; i < Teams; i++) {
      console.log(this.b.TeamsList[i]);
    }

    const shuffledTeams = this.shuffle(this.b.TeamsList);
  
    for (let i = 0; i < shuffledTeams.length; i += 2) {
      const team1 = shuffledTeams[i];
      const team2 = shuffledTeams[i + 1];
      const newMatch = new Match(team1, team2, 0, 0);
      this.b.MatchList.push(newMatch);
    }


    this.bl.push(this.b);
    var t = this.b.Teams;
    this.b.Rounds = Math.ceil(Math.log2(this.b.Teams)) + 1;
    for (let i = 1; i <= this.b.Rounds; i++) {
      if(this.bl.length == this.b.Rounds - 1){
        break;
      }
      t = Math.floor(t/2);
      if(t == 1){
        var bTemp = new Bracket(1);
        bTemp.TeamsList.push("");

        const shuffledTeams = this.shuffle(bTemp.TeamsList);
  
        for (let i = 0; i < shuffledTeams.length+2; i += 2) {
          bTemp.MatchList.push(new Match("","",0,0));
        }
        

        this.bl.push(bTemp);
      }
      else if(t % 2 != 0){
        t = t + 1;
        var bTemp = new Bracket(t);
        for (let i = 1; i <= t; i++){
          bTemp.TeamsList.push("");
        }

        const shuffledTeams = this.shuffle(bTemp.TeamsList);
  
        for (let i = 0; i < shuffledTeams.length+2; i += 2) {
          bTemp.MatchList.push(new Match("","",0,0));
        }

        this.bl.push(bTemp);
      }
      else{
        var bTemp = new Bracket(t);
        for (let i = 1; i <= t; i++){
          bTemp.TeamsList.push("");
        }

        const shuffledTeams = this.shuffle(bTemp.TeamsList);
  
        for (let i = 0; i < shuffledTeams.length+2; i += 2) {
          bTemp.MatchList.push(new Match("","",0,0));
        }

        this.bl.push(bTemp);
      }
    }
  }

  createPairings(){
    for (let i = 0; i < this.b.TeamsList.length; i++) {
      this.b.MatchList.push(new Match(this.b.TeamsList[i], this.b.TeamsList[i+1],0,0));
      i++;
    }
    for (let i = 0; i < this.b.MatchList.length; i++) {
      console.log(this.b.MatchList[i].Member1 + " " + this.b.MatchList[i].Member2)
    }
   
  }

  progressTeam(team: string){
    if(!this.bl[this.bl.length - 1].TeamsList.includes("")){
      for(let i = 0; i < this.bl.length; i++){
        for(let j = 0; j < this.bl[i].MatchList.length; j++){
          if(this.bl[i].MatchList[j].Member1 == ""){
            return
          }
          else if(this.bl[i].MatchList[j].Member2 == ""){
            return
          }
        }
      }
      this.winner = team;
      return;
    }

    if(this.bl[this.currentBracket + 1].TeamsList.includes(team)){
      if(!this.bl[this.currentBracket + 1].TeamsList.includes(""))
      {
        this.currentBracket++;
        for(let i = 0; i < this.bl[this.currentBracket + 1].TeamsList.length; i++){
          if(this.bl[this.currentBracket + 1].TeamsList[i] == "")
          {
            this.bl[this.currentBracket + 1].TeamsList[i] = team;
            return;
          }
        }
      }
      return;
    }
    for(let i = 0; i < this.bl[this.currentBracket + 1].TeamsList.length; i++){
      if(this.bl[this.currentBracket + 1].TeamsList[i] == "")
      {
        this.bl[this.currentBracket + 1].TeamsList[i] = team;
        return;
      }
    }
    this.currentBracket++;

    if(this.currentBracket == this.b.Rounds){

    }

    for(let i = 0; i < this.bl[this.currentBracket + 1].TeamsList.length; i++){
      if(this.bl[this.currentBracket + 1].TeamsList[i] == "")
      {
        this.bl[this.currentBracket + 1].TeamsList[i] = team;
        return;
      }
    }
  }

  updateMatches(team:string){
    var createMatch;
    for(let i = 0; i < this.currentBracket + 2; i++){
      createMatch = true;
      for(let j = 0; j < this.bl[i].MatchList.length; j++){
        if(this.bl[i].MatchList[j].Member1 == team || this.bl[i].MatchList[j].Member2 == team){
          createMatch = false;
          break;
        }
      }
      if(createMatch){
        for(let k = 0; k < this.bl[i].MatchList.length; k++){
          if(this.bl[i].MatchList[k].Member1 == ""){
            this.bl[i].MatchList[k].Member1 = team;
            return;
          }
          else if(this.bl[i].MatchList[k].Member2 == ""){
            this.bl[i].MatchList[k].Member2 = team;
            return;
          }
        }
      }
    }

  }

  shuffle(array:any) {
    for (let i = array.length - 1; i > 0; i--) {
      const j = Math.floor(Math.random() * (i + 1));
      [array[i], array[j]] = [array[j], array[i]];
    }
    return array;
  }

}
