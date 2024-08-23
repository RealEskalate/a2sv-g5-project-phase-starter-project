import 'package:flutter/material.dart';

class TextInputter extends StatefulWidget {
  const TextInputter({Key? key}) : super(key: key);

  @override
  _TextInputterState createState() => _TextInputterState();
}

class _TextInputterState extends State<TextInputter> {
  @override
  Widget build(BuildContext context) {
    return Scaffold(
      body: Center(
        child: Container(
          height: 70,
          width: 500, 
          color: Colors.white,
          child: Row(
            children: [
              const Icon(Icons.attach_file),
              Expanded(
                child: Padding(
                  padding: const EdgeInsets.symmetric(horizontal: 15, vertical: 8),
                  child: TextField(
                    controller: TextEditingController(),
                    decoration: InputDecoration(
                      hintText: "Write your message",
                      hintStyle: const TextStyle(
                        fontFamily: 'Poppins',
                        fontWeight: FontWeight.w400,
                        fontSize: 12,
                        color: Color(0XFF797C7B),
                      ),
                      fillColor: const Color.fromARGB(255, 229, 233, 233),
                      filled: true,
                      border: OutlineInputBorder(
                        borderRadius: BorderRadius.circular(20),
                        borderSide: const BorderSide(
                          width: 0,
                          style: BorderStyle.none,
                        ),
                      ),
                      suffixIcon: Icon(Icons.send), // Correct placement of suffixIcon
                    ),
                  ),
                ),
              ),
              const Icon(Icons.camera_alt_outlined),
              const SizedBox(width: 10),
              const Icon(Icons.keyboard_voice_outlined),
            ],
          ),
        ),
      ),
    );
  }
}
