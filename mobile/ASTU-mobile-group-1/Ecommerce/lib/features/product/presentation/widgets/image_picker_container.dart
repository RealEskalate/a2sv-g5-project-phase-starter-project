import 'dart:io';

import '../../../auth/presentation/pages/pages.dart';

class ImagePickerContainer extends StatefulWidget {
  const ImagePickerContainer({super.key, required this.pickImage});

  final Future<File?> Function() pickImage;

  @override
  State<ImagePickerContainer> createState() => _ImagePickerContainerState();
}

class _ImagePickerContainerState extends State<ImagePickerContainer> {
  File? _image;

  @override
  Widget build(BuildContext context) {
    return GestureDetector(
      onTap: () async {
        final ret = await widget.pickImage();
        if (ret != null) {
          setState(() {
            _image = ret;
          });
        }
      },
      child: Container(
        height: 190,
        width: double.infinity,
        decoration: BoxDecoration(
          color: const Color.fromARGB(255, 230, 230, 230),
          borderRadius: const BorderRadius.all(
            Radius.circular(16),
          ),
          image: _image != null
              ? DecorationImage(
                  image: FileImage(_image!),
                  fit: BoxFit.cover,
                )
              : null,
        ),
        child: _image == null
            ? const Column(
                mainAxisAlignment: MainAxisAlignment.center,
                children: [
                  Icon(
                    Icons.image,
                    size: 36,
                  ),
                  SizedBox(
                    height: 24,
                  ),
                  CustomText(
                    text: 'Upload Image',
                  ),
                ],
              )
            : null,
      ),
    );
  }
}
