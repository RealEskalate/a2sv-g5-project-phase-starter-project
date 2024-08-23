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
      backgroundColor: Colors.white,
      body: Container(
        
        height: double.infinity,
        width: double.infinity,
        color: Colors.white,
        child: Padding(
          padding: const EdgeInsets.symmetric(vertical: 10),
          child: Row(
            mainAxisAlignment: MainAxisAlignment.center,
            children: [
              IconButton(
                onPressed: () {},
                icon: Icon(Icons.attach_file)
                ),
              Expanded(
              
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
                        borderRadius: BorderRadius.circular(25),
                        borderSide: const BorderSide(
                          width: 0,
                          style: BorderStyle.none,
                        ),
                      ),
                      suffixIcon:
                          IconButton(icon:Icon(Icons.send, size: 20),
                          onPressed: (){
                  
                          }, // Correct placement of suffixIcon
                    ),
                  ),),),
                
              const SizedBox(width: 5),
              Row(
              crossAxisAlignment: CrossAxisAlignment.end,
                children: [
          
              IconButton(
                icon: Icon(Icons.camera_alt_outlined),
                onPressed: () {},),
              IconButton(
                icon: Icon(Icons.keyboard_voice_outlined),
                onPressed: () {},),
              ],)
            ],
          ),
        ),
      ),
    );
  }
}
