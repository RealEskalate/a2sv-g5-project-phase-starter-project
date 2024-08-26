import 'dart:ffi';

import 'package:ecommerce_app_ca_tdd/features/chat/presentation/bloc/message_bloc/message_bloc_bloc.dart';
import 'package:ecommerce_app_ca_tdd/features/chat/presentation/bloc/message_bloc/message_bloc_event.dart';
import 'package:flutter/material.dart';
import 'package:flutter_bloc/flutter_bloc.dart';

class ChatBottomAppBar extends StatefulWidget{
  final TextEditingController messageController;
  final VoidCallback onSend;
  final VoidCallback onCamera;
  final VoidCallback onGallery;
  final VoidCallback onAudio;
  final VoidCallback onFile;

  ChatBottomAppBar({
    required this.messageController,
    required this.onCamera,
    required this.onSend,
    required this.onGallery,
    required this.onAudio,
    required this.onFile,
  });

  @override
  _ChatBottomAppBarState createState() => _ChatBottomAppBarState();

}
class _ChatBottomAppBarState extends State<ChatBottomAppBar>{
  
  bool _isTextFieldEmpty = true;
  
  @override
  void initState(){
    super.initState();
    widget.messageController.addListener(_onMessageChanged);
  }

  void dispose(){
    widget.messageController.removeListener(_onMessageChanged);
    super.dispose();
  }

  void _onMessageChanged(){
    setState(() {
      _isTextFieldEmpty = widget.messageController.text.isEmpty;
    });
  }

  @override
  Widget build(BuildContext context) {
     void onSendAction (String chatID,TextEditingController _messageController,String type){
    context.read<MessageBloc>().add(MessageSent(chatID, _messageController.text, 'text'));
    _messageController.clear();
  }
   
    return BottomAppBar(
      shadowColor: Colors.grey,
      elevation: 10,
      color: Colors.white,
      child: Row(
          children: [
            Padding(
              padding: const EdgeInsets.only(left: 0.0),
              child: IconButton(
                icon: Icon(Icons.attach_file),
                onPressed: widget.onFile,
                padding: EdgeInsets.zero, 
              ),
            ),
            Expanded(
              child: TextField(
              controller: widget.messageController,
              decoration: InputDecoration(
                fillColor: Color.fromRGBO(243, 246, 246, 1),
                hintText: 'Type a message',
                border: OutlineInputBorder(
                  borderRadius: BorderRadius.circular(15),
                  borderSide: BorderSide.none,
                ),
                filled: true,
                contentPadding: EdgeInsets.symmetric(horizontal: 20, vertical: 15),
              suffixIcon: IconButton(
                            onPressed: _isTextFieldEmpty ? widget.onGallery : widget.onSend,
                            icon: Icon(_isTextFieldEmpty ? Icons.photo : Icons.send, color: _isTextFieldEmpty ? Colors.blue : Colors.blue,)),
                            ),
              )
            ),
            
            IconButton(onPressed: widget.onCamera, icon: Icon(Icons.camera_alt), padding: EdgeInsets.zero, 
              alignment: Alignment.centerRight,  ),
            IconButton(onPressed: widget.onAudio, icon: Icon(Icons.mic),padding: EdgeInsets.zero, 
              alignment: Alignment.centerRight,  ),
          ],
        ),
        
    
  );
  }

}
