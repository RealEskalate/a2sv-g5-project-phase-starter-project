import 'dart:async';
import 'dart:io';

import 'package:flutter/material.dart';
import 'package:flutter_bloc/flutter_bloc.dart';
import 'package:image_picker/image_picker.dart';

import '../../../../core/themes/themes.dart';
import '../bloc/cubit/input_validation_cubit.dart';

class PickImage extends StatelessWidget {
  const PickImage({
    super.key,
  });
  FutureOr<File?> pickImage() async {
    final imagePicker = ImagePicker();
    final pickFile = await imagePicker.pickImage(source: ImageSource.gallery);
    if (pickFile != null) {
      return File(pickFile.path);
    }
    return null;
  }

  @override
  Widget build(BuildContext context) {
    return BlocBuilder<InputValidationCubit, InputValidationState>(
      builder: (context, state) {
        if (state is InputValidationInitial ||
            (state is InputValidatedState && state.imageUrl == null)) {
          return Stack(
            alignment: Alignment.center,
            children: [
              Container(
                height: 180,
                margin: const EdgeInsets.symmetric(
                  horizontal: 20,
                  vertical: 10,
                ),
                decoration: BoxDecoration(
                  color: MyTheme.ecInputGrey,
                  borderRadius: BorderRadius.circular(10),
                ),
              ),
              Column(
                children: [
                  IconButton(
                    onPressed: () async {
                      File? myImage = await pickImage();
                      if (myImage != null) {
                        // ignore: use_build_context_synchronously
                        BlocProvider.of<InputValidationCubit>(context)
                            .setImage(myImage);
                      }
                    },
                    icon: const Icon(
                      Icons.image_outlined,
                      size: 80,
                    ),
                  ),
                  const Text('Upload Image')
                ],
              ),
            ],
          );
        } else {
          return Container(
            height: 180,
            margin: const EdgeInsets.symmetric(
              horizontal: 20,
              vertical: 10,
            ),
            decoration: BoxDecoration(
              color: MyTheme.ecInputGrey,
              borderRadius: BorderRadius.circular(10),
            ),
            child: ClipRRect(
              borderRadius: BorderRadius.circular(10),
              child: SizedBox(
                height: 180,
                width: double.infinity,
                child: Image.file(state.imageUrl!),
              ),
            ),
          );
        }
      },
    );
  }
}
