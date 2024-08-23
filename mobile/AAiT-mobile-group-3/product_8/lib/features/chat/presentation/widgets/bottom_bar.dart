import 'dart:io';

import 'package:flutter/material.dart';
import 'package:flutter_svg/flutter_svg.dart';
import 'package:image_picker/image_picker.dart';
import 'package:file_picker/file_picker.dart';

class MyMessageBar extends StatefulWidget {
  const MyMessageBar({
    super.key,
  });

  @override
  State<MyMessageBar> createState() => _MyMessageBarState();
}

class _MyMessageBarState extends State<MyMessageBar> {
  final TextEditingController message = TextEditingController();

  File? _image;

  final ImagePicker _picker = ImagePicker();

  Future<void> _pickImage() async {
    final XFile? pickedFile =
        await _picker.pickImage(source: ImageSource.gallery);
    if (pickedFile != null) {
      setState(() {
        _image = File(pickedFile.path); // Store the selected image in _image
      });
    }
  }

  File? _file;

  Future<void> _pickFile() async {
    FilePickerResult? result = await FilePicker.platform.pickFiles();

    if (result != null && result.files.single.path != null) {
      setState(() {
        _file =
            File(result.files.single.path!); // Store the selected file in _file
      });
    }
  }

  var copied = '';

  @override
  Widget build(BuildContext context) {
    return Scaffold(
        body: Positioned(
      child: Center(
        child: Card(
          shadowColor: Colors.grey,
          child: SingleChildScrollView(
            child: Row(children: [
              IconButton(
                icon: SvgPicture.asset('assets/icons/clip.svg'),
                onPressed: () {
                  _pickFile();
                },
              ),
              Expanded(
                child: Padding(
                  padding: const EdgeInsets.symmetric(horizontal: 1),
                  child: TextField(
                    controller: message,
                    decoration: InputDecoration(
                        border: OutlineInputBorder(
                            borderRadius: BorderRadius.circular(30),
                            borderSide: BorderSide.none),
                        filled: true,
                        fillColor: const Color.fromARGB(255, 243, 246, 246),
                        hintText: 'Write your message here',
                        hintStyle: const TextStyle(
                            fontSize: 15, fontWeight: FontWeight.w100),
                        suffixIcon: IconButton(
                          icon: SvgPicture.asset(
                            'assets/icons/copy.svg',
                            width: 20,
                            height: 20,
                          ),
                          onPressed: () {
                            print('copied');
                            copied = message.text;
                          },
                        )),
                  ),
                ),
              ),
              IconButton(
                icon: SvgPicture.asset('assets/icons/camera.svg'),
                onPressed: () {
                  print('pressed');
                  _pickImage();
                },
              ),
              IconButton(
                icon: SvgPicture.asset('assets/icons/microphone.svg'),
                onPressed: () {
                  print('pressed');
                },
              ),
              IconButton(
                icon: SvgPicture.asset('assets/icons/send.svg'),
                onPressed: () {
                  print('pressed');
                },
              ),
            ]),
          ),
          // ),
        ),
      ),
    ));
  }
}
