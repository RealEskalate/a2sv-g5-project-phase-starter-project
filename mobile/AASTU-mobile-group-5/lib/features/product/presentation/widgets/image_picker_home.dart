import 'dart:io';
import 'package:flutter/material.dart';
import 'package:image_picker/image_picker.dart';

class ImagePickerHome extends StatefulWidget {
  final double containerHeight;
  final double containerWidth;
  final BoxDecoration decoration;

  const ImagePickerHome({
    super.key,
    required this.containerHeight,
    required this.containerWidth,
    required this.decoration,
  });

  @override
  // ignore: library_private_types_in_public_api
  _ImagePickerHomeState createState() => _ImagePickerHomeState();
}

class _ImagePickerHomeState extends State<ImagePickerHome> {
  XFile? _image;

  Future<void> _pickImage(ImageSource source) async {
    final ImagePicker picker = ImagePicker();
    final XFile? image = await picker.pickImage(source: source);

    if (image != null) {
      setState(() {
        _image = image;
      });
    }
  }

  @override
  Widget build(BuildContext context) {
    return Container(
      decoration: widget.decoration,
      height: widget.containerHeight,
      width: widget.containerWidth,
      child: Center(
        child: _image == null
            ? IconButton(
                icon: Icon(Icons.add_a_photo, color: Colors.grey[600]),
                onPressed: () => _showImageSourceDialog(context),
              )
            : ClipRRect(
                borderRadius: BorderRadius.circular(10),
                child: Image.file(
                  File(_image!.path),
                  fit: BoxFit.cover,
                  width: widget.containerWidth,
                  height: widget.containerHeight,
                ),
              ),
      ),
    );
  }

  void _showImageSourceDialog(BuildContext context) {
    showDialog(
      context: context,
      builder: (BuildContext context) {
        return AlertDialog(
          title: const Center(child: Text('Pick an image')),
          
          actions: <Widget>[
            Center(
              child: Row(
                mainAxisSize: MainAxisSize.min,
                children: [
                  TextButton(
                    child: const Text('Camera', style: TextStyle(color: Colors.blue)),
                    onPressed: () {
                      Navigator.of(context).pop();
                      _pickImage(ImageSource.camera);
                    },
                  ),
                  TextButton(
                    child: const Text('Gallery', style: TextStyle(color: Colors.blue)),
                    onPressed: () {
                      Navigator.of(context).pop();
                      _pickImage(ImageSource.gallery);
                    },
                  ),
                ],
              ),
            ),
          ],
        );
      },
    );
  }
}

// import 'dart:io';

// import 'package:flutter/material.dart';
// import 'package:image_picker/image_picker.dart';

// class ImagePickerHome extends StatefulWidget {
//   final double containerHeight;
//   final double containerWidth;
//   final BoxDecoration decoration;

//   const ImagePickerHome({
//     Key? key,
//     required this.containerHeight,
//     required this.containerWidth,
//     required this.decoration,
//   }) : super(key: key);

//   @override
//   _ImagePickerHomeState createState() => _ImagePickerHomeState();
// }

// class _ImagePickerHomeState extends State<ImagePickerHome> {
//   XFile? _image;

//   Future<void> _pickImage(ImageSource source) async {
//     final ImagePicker picker = ImagePicker();
//     final XFile? image = await picker.pickImage(source: source);

//     if (image != null) {
//       setState(() {
//         _image = image;
//       });
//     }
//   }

//   @override
//   Widget build(BuildContext context) {
//     return Container(
//       decoration: widget.decoration,
//       height: widget.containerHeight,
//       width: widget.containerWidth,
//       child: Center(
//         child: _image == null
//             ? IconButton(
//                 icon: Icon(Icons.add_a_photo, color: Colors.grey[600]),
//                 onPressed: () => _showImageSourceDialog(context),
//               )
//             : ClipRRect(
//                 borderRadius: BorderRadius.circular(10),
//                 child: Image.file(
//                   File(_image!.path),
//                   fit: BoxFit.cover,
//                   width: widget.containerWidth,
//                   height: widget.containerHeight,
//                 ),
//               ),
//       ),
//     );
//   }

//   void _showImageSourceDialog(BuildContext context) {
//     showDialog(
//       context: context,
//       builder: (BuildContext context) {
//         return AlertDialog(
//           title: const Text('Pick an image'),
//           actions: <Widget>[
//             TextButton(
//               child: const Text('Camera'),
//               onPressed: () {
//                 Navigator.of(context).pop();
//                 _pickImage(ImageSource.camera);
//               },
//             ),
//             TextButton(
//               child: const Text('Gallery'),
//               onPressed: () {
//                 Navigator.of(context).pop();
//                 _pickImage(ImageSource.gallery);
//               },
//             ),
//           ],
//         );
//       },
//     );
//   }
// }