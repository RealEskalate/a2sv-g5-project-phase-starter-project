import 'package:ecommerce/features/chat_feature/chat/data_layer/data_source/Service/socker_service.dart';
import 'package:ecommerce/features/chat_feature/chat/domain/usecase/delete_chat.dart';
import 'package:ecommerce/features/chat_feature/chat/domain/usecase/get_messages.dart';
import 'package:flutter/cupertino.dart';
import 'package:flutter/material.dart';
import 'package:google_fonts/google_fonts.dart';
import 'package:image_picker/image_picker.dart';
import '../../domain/usecase/send_message.dart';
import '../widget/left_chat.dart';
import '../widget/right_chat.dart';
import 'data.dart';
import 'package:socket_io_client/socket_io_client.dart' as IO;

// ignore: must_be_immutable
class Directmessage extends StatefulWidget {
  SendMessage sendMessage;
  GetMessages getMessages;
  Directmessage({required this.getMessages,required this.sendMessage, super.key});

  @override
  State<Directmessage> createState() => _DirectmessageState();
}

class _DirectmessageState extends State<Directmessage> {
  TextEditingController messageController = TextEditingController();
  final ImagePicker _picker = ImagePicker();
  String imagePath = '';
  String audioPath = '';

  bool isRecording = false;
  bool isPlaying = false;
  String chatId = "66c82e01bb16bfe67cd3b541";

  Future<void> _pickImage(ImageSource source) async {
    final pickedFile = await _picker.pickImage(source: source);
    if (pickedFile != null) {
      setState(() {
        imagePath = pickedFile.path;
      });
    } else {
      debugPrint('No image selected.');
    }
  }

  @override
  void initState() {
    super.initState();
  }

  @override
  void dispose() {
    // TODO: implement dispose

    super.dispose();
  }

  @override
  Widget build(BuildContext context) {
    String id = "001";
    double width = MediaQuery.of(context).size.width;
    return Scaffold(
      backgroundColor: Colors.white,
      body: Column(
        children: [
          // Header Row
          Container(
            padding: const EdgeInsets.only(top: 25, left: 20, right: 20),
            child:  Row(
              children: [
                Padding(
                  padding: EdgeInsets.only(right: 20),
                  child: Icon(Icons.arrow_back),
                ),
                CircleAvatar(
                  backgroundColor: Color(0XFFFEC7D3),
                  backgroundImage: AssetImage('assets/images/splash.png'),
                  radius: 30,
                ),
                SizedBox(width: 20),
                Text(
                  'Username',
                ),
                Spacer(),
                IconButton(
                  onPressed: 
                  (){
                      widget.getMessages.call(chatId);
                  }
                  ,
                  icon:Icon(CupertinoIcons.phone,size: 25,),
                  
                ),
                SizedBox(width: 8),
                Icon(
                  CupertinoIcons.video_camera,
                  size: 30,
                ),
              ],
            ),
          ),

          // Messages List
          Expanded(
            child: ListView.builder(
              reverse: true,
              padding: const EdgeInsets.all(10), // Add padding if needed
              itemCount: messages.length,
              itemBuilder: (context, index) {
                if (messages[index].sender.id == id) {
                  return RightChat(message: messages[index]);
                } else {
                  return LeftChat(message: messages[index]);
                }
              },
            ),
          ),

          // Input Row
          Container(
            padding: const EdgeInsets.all(10),
            color: Colors
                .white, // Ensure the background color matches or is suitable
            child: Row(
              crossAxisAlignment: CrossAxisAlignment.end,
              children: [
                Transform.rotate(
                  angle: 45 * (3.1415926535897932 / 180),
                  child: IconButton(
                    icon: const Icon(Icons.attach_file),
                    onPressed: () {
                      _pickImage(ImageSource.gallery);
                    },
                  ),
                ),
                Expanded(
                  child: TextField(
                    controller: messageController,
                    decoration: InputDecoration(
                      hintText: 'message',
                      hintStyle: const TextStyle(
                        color: Color(0xFF797C7B),
                      ),
                      suffixIcon: messageController.text.isEmpty
                          ? IconButton(
                              onPressed: () {
                              
                                widget.sendMessage.call(
                                    chatId, messageController.text, 'text');
                              },
                              icon: const Icon(CupertinoIcons.square_on_square),
                            )
                          : IconButton(
                              onPressed: () {
                                widget.sendMessage.call(
                                    chatId, messageController.text, 'text');
                              },
                              icon: const Icon(
                                Icons.send,
                                size: 25,
                              )),
                      border: OutlineInputBorder(
                        borderRadius: BorderRadius.circular(15),
                        borderSide: BorderSide.none,
                      ),
                      filled: true,
                      fillColor: const Color(0xFFF3F6F6),
                    ),
                  ),
                ),
                if (messageController.text.isEmpty)
                  IconButton(
                    icon: const Icon(CupertinoIcons.camera),
                    onPressed: () {
                      _pickImage(ImageSource.camera);
                    },
                  ),
                if (messageController.text.isEmpty)
                  IconButton(
                    icon: isRecording
                        ? const Icon(CupertinoIcons.stop_circle,
                            color: Colors.red)
                        : const Icon(CupertinoIcons.mic),
                    onPressed: () {},
                  ),
              ],
            ),
          ),
        ],
      ),
    );
  }
}

var messages = [];
