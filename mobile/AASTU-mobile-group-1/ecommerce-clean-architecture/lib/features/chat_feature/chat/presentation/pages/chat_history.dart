import 'package:ecommerce/features/chat_feature/chat/presentation/widget/avatar.dart';
import 'package:ecommerce/features/chat_feature/chat/presentation/widget/chat_card.dart';
import 'package:flutter/material.dart';
import 'package:flutter/widgets.dart';
import 'package:google_fonts/google_fonts.dart';

class UiChatHistory extends StatefulWidget {
  const UiChatHistory({super.key});

  @override
  State<UiChatHistory> createState() => _UiChatHistoryState();
}

class _UiChatHistoryState extends State<UiChatHistory> {
  @override
  Widget build(BuildContext context) {
    return Scaffold(
      backgroundColor: Color(0xFF498cf0),
      body: Padding(
        padding: EdgeInsets.only(top: 30,),
          child: Column(
            children: [
              Row(
                crossAxisAlignment: CrossAxisAlignment.start,
                children: [
                  IconButton(
                    onPressed:(){

                    }, 
                    icon: Icon(Icons.search),
                    color: Colors.white,
                    iconSize : 40,
                    )
              ],),
              SizedBox(height: 10,),
              SingleChildScrollView(
                scrollDirection: Axis.horizontal,
                child: Row(
                  children: [
                  SizedBox(width: 15),
                  avatar(name:'Name'),
                  SizedBox(width: 15),
                  avatar(name:'Name'),
                  SizedBox(width: 15),
                  avatar(name:'Name'),
                  SizedBox(width: 15),
                  avatar(name:'Name'),
                  SizedBox(width: 15),
                  avatar(name:'Name'),
                  SizedBox(width: 15),
                  
                  ],),
              ),
            SizedBox(height: 60,),
            Expanded(
              child: Container(
                width: double.infinity,
                decoration: BoxDecoration(
                  color: Colors.white,
                  borderRadius: BorderRadius.only(
                    topLeft: Radius.circular(50),
                    topRight: Radius.circular(50),
                  ),
                ),
              child: SingleChildScrollView(
                child: Column(
                  crossAxisAlignment: CrossAxisAlignment.center,
                  children: [
                    SizedBox(height: 15,),
                    Container(
                    
                      width:40,
                      height:7,
                      decoration: BoxDecoration(
                        borderRadius: BorderRadius.circular(30),
                        color:Color(0xFFE6E6E6)
                        )
                      
                    ),
                    SizedBox(height: 20),
                    ChatCard(
                      name: 'Name',
                      message: 'Message',
                      time: 'Time',
                      imagePath: 'assets/images/avatar.jpg',
                      onDelete: (){},
                    ),
                    SizedBox(height: 20),
                    ChatCard(
                
                      name: 'Name',
                      message: 'Message',
                      time: 'Time',
                      imagePath: 'assets/images/avatar.jpg',
                      onDelete: (){},
                    ),


                    SizedBox(height: 20),
                    ChatCard(
                    
                      name: 'Name',
                      message: 'Message',
                      time: 'Time',
                      imagePath: 'assets/images/avatar.jpg',
                      onDelete: (){},
                    ),
                    SizedBox(height: 20),
                    ChatCard(
                     
                      name: 'Name',
                      message: 'Message',
                      time: 'Time',
                      imagePath: 'assets/images/avatar.jpg',
                      onDelete: (){},
                    ),
                    SizedBox(height: 20),
                    ChatCard(
                    
                      name: 'Name',
                      message: 'Message',
                      time: 'Time',
                      imagePath: 'assets/images/avatar.jpg',
                      onDelete: (){},
                    ),
                    SizedBox(height: 20),
                    ChatCard(
                      
                      name: 'Name',
                      message: 'Message',
                      time: 'Time',
                      imagePath: 'assets/images/avatar.jpg',
                      onDelete: (){},
                    ),
                    SizedBox(height: 20),
                    ChatCard(
                
                      name: 'Name',
                      message: 'Message',
                      time: 'Time',
                      imagePath: 'assets/images/avatar.jpg',
                      onDelete: (){},
                    )
             
                  ],
                ),
              )
              ),
            ),
            ],
          )
          
        ),
      );
  }
}

