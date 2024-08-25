import 'package:file_picker/file_picker.dart';
import 'package:flutter/material.dart';
import 'package:image_picker/image_picker.dart';

class TextInputter extends StatefulWidget {
  const TextInputter({Key? key}) : super(key: key);

  @override
  _TextInputterState createState() => _TextInputterState();
}

class _TextInputterState extends State<TextInputter> {
  final TextEditingController _controller = TextEditingController();
  bool isRecording = false;
  bool isTyping = false;
  final ImagePicker _picker = ImagePicker();

   @override
  void initState() {
    super.initState();

    _controller.addListener(() {
      setState(() {
        isTyping = _controller.text.isNotEmpty;
      });
    });
  }

  @override
  void dispose() {
    _controller.dispose();
    super.dispose();
  }

  void _pickFile() async {
    FilePickerResult? result = await FilePicker.platform.pickFiles();
    if (result != null) {
      // Handle file selection
    }
  }

  void _pickImageFromGallery() async {
    final pickedFile = await _picker.pickImage(source: ImageSource.gallery);
    if (pickedFile != null) {
      // Handle image selection
    }
  }

  void _openCamera() async {
    final pickedFile = await _picker.pickImage(source: ImageSource.camera);
    if (pickedFile != null) {
      // Handle image capture
    }
  }

  void _toggleRecording() {
    setState(() {
      isRecording = !isRecording;
      // Add your recording logic here
    });
  }

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
            // mainAxisAlignment: MainAxisAlignment.center,
            children: [
              IconButton(
                onPressed: _pickFile,
                icon: Icon(Icons.attach_file)
                ),
              Expanded(
              
                  child: TextField(
                    controller: _controller,
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
                      suffixIcon: IconButton(
                        icon:Icon(
                            isTyping ? Icons.send:
                            Icons.image,
                            size: 20, 
                            color : const Color.fromARGB(255, 95, 94, 94)
                            ),
                            onPressed: isTyping ? (){} : _pickImageFromGallery, // Correct placement of suffixIcon
                    ),
                  ),),),
                
              const SizedBox(width: 5),
              Row(
              mainAxisAlignment: MainAxisAlignment.end,
                children: [
          
              IconButton(
                icon: Icon(Icons.camera_alt),
                onPressed: _openCamera),
              IconButton(
                icon: Icon(isRecording ? Icons.stop : Icons.mic,  ),
                onPressed: _toggleRecording,),
              ],)
            ],
          ),
        ),
      ),
    );
  }
}
