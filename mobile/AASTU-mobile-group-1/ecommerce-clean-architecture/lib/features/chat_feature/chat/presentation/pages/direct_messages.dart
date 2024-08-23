import 'package:flutter/material.dart';
import 'package:google_fonts/google_fonts.dart';

import '../widget/left_chat.dart';
import '../widget/right_chat.dart';
import 'data.dart';

class Directmessage extends StatelessWidget {
  const Directmessage({super.key});

  @override
  Widget build(BuildContext context) {
    String id = "001";

    return Scaffold(
      backgroundColor: Colors.white,
      body: Column(
        children: [
          // Header Row
          Container(
            padding: const EdgeInsets.only(top: 25, left: 20, right: 20),
            child: Row(
              children: [
                const Padding(
                  padding: EdgeInsets.only(right: 20),
                  child: Icon(Icons.arrow_back),
                ),
                const CircleAvatar(
                  backgroundColor: Color(0XFFFEC7D3),
                  backgroundImage: AssetImage('assets/images/app_icon.png'),
                  radius: 35,
                ),
                const SizedBox(width: 20),
                Text(
                  'Username',
                  style: GoogleFonts.poppins(
                    fontWeight: FontWeight.w500,
                    fontSize: 18,
                  ),
                ),
                const Spacer(),
                const Icon(Icons.phone_outlined),
                const SizedBox(width: 5),
                const Icon(Icons.video_call),
              ],
            ),
          ),

          // Messages List
          Expanded(
            child: ListView.builder(
              reverse: true,
              padding: EdgeInsets.all(10), // Add padding if needed
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
                      // Handle the button press
                    },
                  ),
                ),
                Expanded(
                  child: TextField(
                    decoration: InputDecoration(
                      hintText: "message",
                      hintStyle: const TextStyle(
                        color: Color(0xFF797C7B),
                      ),
                      suffixIcon: IconButton(
                        onPressed: () {},
                        icon: const Icon(Icons.content_copy_outlined),
                      ),
                      border: OutlineInputBorder(
                        borderRadius: BorderRadius.circular(15),
                        borderSide: BorderSide.none,
                      ),
                      filled: true,
                      fillColor: const Color(0xFFF3F6F6),
                    ),
                  ),
                ),
                IconButton(
                  icon: const Icon(Icons.camera_alt_outlined),
                  onPressed: () {},
                ),
                IconButton(
                  icon: const Icon(Icons.mic_none),
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
