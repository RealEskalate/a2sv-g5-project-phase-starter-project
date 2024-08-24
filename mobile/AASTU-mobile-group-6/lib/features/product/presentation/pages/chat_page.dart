import 'package:flutter/material.dart';
import 'package:temp_project/chat_bottom_appbar.dart';
import 'package:temp_project/chat_appbar.dart';
import 'package:temp_project/chat_body.dart';

class ChatPage extends StatelessWidget{
  final TextEditingController _messageController = TextEditingController();

  // Dummy data for chat
  final String ownerId = '1'; 
  final List<String> chats = [
    'Hey, how are you?',
    'I\'m good, thanks! How about you?',
    'Doing well, just working on some projects.',
    'That sounds interesting! What kind of projects?',
    'I\'m working on a Flutter app, it\'s quite fun!',
    'That\'s awesome! Flutter is really powerful.',
    'Yeah, I love how easy it is to build UI with Flutter.',
    'I totally agree! What feature are you working on?',
    'I\'m implementing a chat feature, just like this one!',
    'Nice! I can\'t wait to see the final product.',
  ];
  
  final List<String> senderIds = [
    '2', // Sender 1
    '1', // Owner
    '2',
    '1',
    '2',
    '1',
    '2',
    '1',
    '2',
    '1',
  ];


  @override
  Widget build(BuildContext context){
    return MaterialApp(
      home: Scaffold(
        appBar:
          ChatAppBar('Meron Weldu', 'last seen yesterday', "lib/temp_photo.jpg", isOnline: true),
       
        body: Container(
          color: Colors.white,
          child: Column(
            children: [
              // ChatBody to display the chat messages
              Expanded(
                child: ChatBody(
                  ownerId: ownerId,
                  senderIds: senderIds,
                  chats: chats,
                ),
              ),

              ChatBottomAppBar( 
                messageController: _messageController,
                onSend: () {
                  print('Send button pressed');
                },
                onFile: () {
                  print('File button pressed');
                }, 
                onCamera: () {
                  print('Camera button pressed');
                },
                onGallery: () {
                  print('Gallery button pressed');
                },
                onAudio: () {
                  print('Audio button pressed');
                },
              ),
            ],
          ),
        )
      ),
    );
  }
 }
