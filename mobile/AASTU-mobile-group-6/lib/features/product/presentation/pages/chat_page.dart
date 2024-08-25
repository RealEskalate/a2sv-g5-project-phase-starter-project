import 'package:ecommerce_app_ca_tdd/features/chat/domain/entities/chat_entity.dart';
import 'package:ecommerce_app_ca_tdd/features/chat/domain/entities/message.dart';
import 'package:ecommerce_app_ca_tdd/features/chat/socket_n/chatbox.dart';
import 'package:ecommerce_app_ca_tdd/features/product/data/models/seller_model.dart';
import 'package:ecommerce_app_ca_tdd/features/product/presentation/widgets/chat_appbar.dart';
import 'package:ecommerce_app_ca_tdd/features/product/presentation/widgets/chat_bottom_appbar.dart';
import 'package:flutter/material.dart';
import 'package:shared_preferences/shared_preferences.dart';
class ChatPage extends StatelessWidget{
  final SellerModel sellerID;
  final TextEditingController _messageController = TextEditingController();

  ChatPage({super.key, required this.sellerID});
  @override
  Widget build(BuildContext context){
    return MaterialApp(
      debugShowCheckedModeBanner: false,
      home: Scaffold(
        appBar:
          ChatAppBar(sellerID.name, 'last seen yesterday', "image.jpg"),
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
                  var chat = Message(
                    senderId: getSenderId().toString(),
                    chatId: '',
                    content: _messageController.text,
                    type: 'text',
                  ); 
                var socket = SocketService().connectToServer();
                
                  if (_messageController.text.isNotEmpty) {
                    SocketService().sendMessage(chat);
                    _messageController.clear();
                  }
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
Future<String> getSenderId() async {
  var prefs = await SharedPreferences.getInstance();
  return prefs.getString('id')!;
}