package progress

import (
	"github.com/ChaitanyaAkula/gittyjobsdb"
	"fmt"

	"log"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)
/*
Image =10
About=20(4 fields * 5 points)
Employment =10(>1)
Projects(experience)= 10(>1)
Education =20(=>3)(for <3 {
 5 fields * 1.33 points = 6.65
})
skills=20(>10)(if <10 each skill =2 points)
*/

func GetProgress(gittyid string) float64 {

	var totalprogress,aboutprogress,expprogress,eduprogress,skillprogress,certprogress,imgprogress,empprogress =0.0,0.0,0.0,0.0,0.0,0.0,0.0,0.0
    db:=dbconnection.Connection()
	defer db.Close()
	var Devimagepath,Devlocation,Devsalary,Devemploymenttype,Devaboutme string
	//var Devempid,Devcompanyname,Devempstartyear,Devempstartmonth,Devempendyear,Devempendmonth,Devemptechnologies,DevPresentworking,Devjoblocation string
   // var Devprojectid,Devprojecttitle,Devprojeturl,Devprojectdescription string
	//var Devlanguages,Devframeworks,Devlibraries,Devotherskills string
	var Devlanguages string
    var Deveducollegename,Devedudegree,Devedulocation,Devedustartyear,Deveduendyear string


	rows1,_:=db.Query("select imagepath from gittyimages WHERE gittyaccountid=? ",gittyid)
			for rows1.Next(){
				err:=rows1.Scan(&Devimagepath)
					
				if err!=nil{
					log.Fatalln(err)
				}
				
	}
	if Devimagepath!="user-icon.png"{
		imgprogress+=10
	}
	rows2,_:=db.Query("select location,aboutme,salary,employmenttype from people WHERE gittyaccountid=? ",gittyid)
			for rows2.Next(){
				err:=rows2.Scan(&Devlocation,&Devaboutme,&Devsalary,&Devemploymenttype)
					
				if err!=nil{
					log.Fatalln(err)
				}
				
	}
	if Devlocation!=""{
		aboutprogress+=5
	}
	if Devemploymenttype!=""{
		aboutprogress+=5
	}
	if Devsalary!=""{
		aboutprogress+=5
	}
	if Devaboutme!=""{
		aboutprogress+=5
	}
/*	rows3,_:=db.Query("select iddevemployment,companyname,jobtitle,startyear,startmonth,endyear,endmonth,technologies,jobdescription,presentcheck,joblocation from devemployment WHERE gittyaccountid=? ",gittyid)
			for row s3.Next(){
				err:=rows3.Scan(&Devempid,&Devcompanyname,&Jobtitle,&Devempstartyear,&Devempstartmonth,&Devempendyear,&Devempendmonth,&Devemptechnologies,&Jobdescription,&DevPresentworking,&Devjoblocation)
					
				if err!=nil{
					log.Fatalln(err)
				}
				
	}*/
	var EmpCount int
	result,_:= db.Query("select count(*) from devemployment WHERE gittyaccountid=? ",gittyid)
	for result.Next() {
		err:= result.Scan(&EmpCount)
		if err != nil {
			log.Fatal(err)
		}
	}
	if EmpCount>=1{
		empprogress+=10
	}
/*	rows4,_:=db.Query("select iddevprojects,projecttitle,projecturl,projectdescription from devprojects WHERE gittyaccountid=? ",gittyid)
			for rows4.Next(){
				err:=rows4.Scan(&Devprojectid,&Devprojecttitle,&Devprojeturl,&Devprojectdescription)
					
				if err!=nil{
					log.Fatalln(err)
				}
				
			
	}*/

	ExpCount:=0
	result4,_:=db.Query("select count(*) from devprojects WHERE gittyaccountid=? ",gittyid)
	
	for result4.Next(){
		err4:=result4.Scan(&ExpCount)
		if err4!=nil{
			log.Fatal(err4)
		}
	}
	if ExpCount>=1{
		expprogress+=10
	}
	/*rows5,_:=db.Query("select languages,frameworks,libraries,other from devskills WHERE gittyaccountid=? ",gittyid)
			for rows5.Next(){
				err:=rows5.Scan(&Devlanguages,&Devframeworks,&Devlibraries,&Devotherskills)
					
				if err!=nil{
					log.Fatalln(err)
				}
				
				
	}*/
	SkillCount :=0
	result5,err5:=db.Query("select languages from devskills WHERE gittyaccountid=? ",gittyid)
	if err5!=nil{
		log.Fatal(err5)
	}	
	for result5.Next(){
		err5=result5.Scan(&Devlanguages)
		if err5!=nil{
			log.Fatal(err5)
		}
	}
	fmt.Println(Devlanguages)
	str:=strings.Split(Devlanguages,",")
	SkillCount=len(str)
	if SkillCount<10 && SkillCount>0{
	for i:=len(str);i>0;i--{
		skillprogress+=2
	}
}
	if SkillCount>=10{
		skillprogress+=20
	}


var Count int
	result8,_:=db.Query("select count(*) from deveducation WHERE gittyaccountid=? ",gittyid)
	
	for result8.Next(){
		err8:=result8.Scan(&Count)
		if err8!=nil{
			log.Fatalln(err8)
		}
	}
	if Count<3{
	
		rows6,_:=db.Query("select collegename,degree,location,startyear,endyear from deveducation WHERE gittyaccountid=? ",gittyid)
		for rows6.Next(){
		  err:=rows6.Scan(&Deveducollegename,&Devedudegree,&Devedulocation,&Devedustartyear,&Deveduendyear)
		  if err!=nil{
			  log.Fatalln(err)
		  }

		if Deveducollegename!=""{
			eduprogress+=1.33
		}
		if Devedudegree!=""{
			eduprogress+=1.33
		}
		if Devedulocation!=""{
			eduprogress+=1.33
		}
		if Devedustartyear!=""{
			eduprogress+=1.33
		}
		if Deveduendyear!=""{
			eduprogress+=1.33
		}
	}
	
}
	if Count>=3{
		eduprogress+=20
		fmt.Println(eduprogress)
	}
/*	result9,err9:=db.Query("select count(*) from certifications where userid=?",cookieid)
	if err9!=nil{
		log.Fatal(err9)
	}
	for result9.Next(){
		err9=result9.Scan(&Count)
		if err9!=nil{
			log.Fatal(err9)
		}
	}
	if Count<=2{
	result7,err7:=db.Query("select certificationid,certificationname,certificationcompany from certifications where userid=?",cookieid)
	if err7!=nil{
		log.Fatal(err7)
	}	
	for result7.Next(){
		err7=result7.Scan(&Certificationid,&Certificationname,&Certificationcompany)
		if err7!=nil{
			log.Fatal(err7)
		}
	}
	if Certificationid!=""{
		certprogress+=1.66
	}
	if Certificationname!=""{
		certprogress+=1.66
	}
	if Certificationcompany!=""{
		certprogress+=1.66
	}
}
	if Count>2{
		certprogress+=10
	}*/
	totalprogress=aboutprogress+expprogress+eduprogress+skillprogress+certprogress+imgprogress+empprogress
	return totalprogress
}
