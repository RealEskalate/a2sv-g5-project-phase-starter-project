import 'package:ecommerce_app_ca_tdd/features/product/presentation/widgets/chat_appbar.dart';
import 'package:ecommerce_app_ca_tdd/features/product/presentation/widgets/chat_bottom_appbar.dart';
import 'package:flutter/material.dart';
class ChatPage extends StatelessWidget{
  final TextEditingController _messageController = TextEditingController();
  @override
  Widget build(BuildContext context){
    return MaterialApp(
      home: Scaffold(
        appBar:
          ChatAppBar('Meron Weldu', 'last seen yesterday', "image.jpg"),
       
        body: Container(
          color: Colors.white,
          child: Column(
            children: [
              Expanded(
                child: ListView.builder(
                  itemCount: 10,
                  itemBuilder: (context, index){
                    return ListTile(
                      title: Text('Message $index'),
                    );
                  },
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
