import 'dart:io';

import 'package:flutter/material.dart';
import 'package:google_fonts/google_fonts.dart';
import 'package:image_picker/image_picker.dart';

class ImageUploadWidget extends StatefulWidget {
  final File? imageFile;
  final void Function(File) onImagePicked;

  const ImageUploadWidget({
    super.key,
    required this.onImagePicked,
    this.imageFile,
  });

  @override
  // ignore: library_private_types_in_public_api
  _ImageUploadWidgetState createState() => _ImageUploadWidgetState();
}

class _ImageUploadWidgetState extends State<ImageUploadWidget> {
  Future<void> _pickImage(BuildContext context, ImageSource source) async {
    final picker = ImagePicker();
    final pickedFile = await picker.pickImage(source: source, imageQuality: 80);

    if (pickedFile != null) {
      File imageFile = File(pickedFile.path);
      widget.onImagePicked(imageFile);
      setState(() {});
    }
  }

  @override
  Widget build(BuildContext context) {
    return GestureDetector(
      onTap: () => _showPicker(context),
      child: Container(
        height: 300,
        width: 400,
        decoration: BoxDecoration(
          color: Theme.of(context).brightness == Brightness.dark ? const Color.fromARGB(255, 138, 134, 134) : const Color.fromRGBO(243, 243, 243, 1),
          borderRadius: BorderRadius.circular(10),
          image: widget.imageFile != null
              ? DecorationImage(
                  image: FileImage(widget.imageFile!),
                  fit: BoxFit.cover,
                )
              : null,
        ),
        padding: const EdgeInsets.all(8.0),
        child: widget.imageFile == null
            ? Center(
                child: Column(
                  mainAxisAlignment: MainAxisAlignment.center,
                  children: [
                    const Icon(Icons.image, color: Colors.black, size: 30),
                    const SizedBox(height: 10),
                    Text(
                      'Upload image',
                      style: GoogleFonts.poppins(
                      color: Theme.of(context).brightness == Brightness.dark ? Colors.white : Colors.black,
                    fontSize: 18,
                    ),
                    ),
                  ],
                ),
              )
            : null,
      ),
    );
  }

  void _showPicker(BuildContext context) {
    showModalBottomSheet(
      context: context,
      builder: (BuildContext bc) {
        return SafeArea(
          child: Wrap(
            children: <Widget>[
              ListTile(
                leading: const Icon(Icons.photo_library),
                title: const Text('Select from Gallery'),
                onTap: () {
                  _pickImage(context, ImageSource.gallery);
                  Navigator.of(context).pop();
                },
              ),
              ListTile(
                leading: const Icon(Icons.photo_camera),
                title: const Text('Take a Photo'),
                onTap: () {
                  _pickImage(context, ImageSource.camera);
                  Navigator.of(context).pop();
                },
              ),
            ],
          ),
        );
      },
    );
  }
}
