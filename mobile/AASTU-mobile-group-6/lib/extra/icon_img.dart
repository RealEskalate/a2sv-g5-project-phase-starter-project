import 'package:flutter/material.dart';
import 'package:image_picker/image_picker.dart';
import 'dart:io';

class ImagePickerIconButton extends StatefulWidget {
  @override
  _ImagePickerIconButtonState createState() => _ImagePickerIconButtonState();
}

class _ImagePickerIconButtonState extends State<ImagePickerIconButton> {
  File? _image;

  Future<void> _pickImage() async {
    final ImagePicker _picker = ImagePicker();
    final XFile? pickedFile =
        await _picker.pickImage(source: ImageSource.gallery);

    if (pickedFile != null) {
      setState(() {
        _image = File(pickedFile.path);
      });
    }
  }

  @override
  Widget build(BuildContext context) {
    return IconButton(
      icon: _image == null
          ? Icon(
              Icons.circle_rounded,
              size: 60.0, 
              color: Color.fromARGB(255, 204, 204, 204),
            )
          : ClipOval(
              child: Image.file(
                _image!,
                width: 50.0,  
                height: 50.0,
                fit: BoxFit.cover, 
              ),
            ),      
        onPressed: _pickImage,
        // iconSize: 80.0,
    );
  }
}
