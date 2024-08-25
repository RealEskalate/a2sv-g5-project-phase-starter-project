
import 'dart:io';

import 'package:equatable/equatable.dart';

abstract class ImageState extends Equatable{}

// ignore: must_be_immutable
class OnImageSelect extends ImageState {
  final String image;
   File? file; // Ensure this can be null

  OnImageSelect({required this.image, this.file});

  @override
  List<Object?> get props => [image];
}

class InputIntialState extends ImageState {
  @override
  List<Object?> get props => [];
}


class ImageLoadingState extends ImageState {
  @override
  List<Object?> get props => [];
}

class ErrorState extends ImageState {
  final String messages;

  ErrorState ({
    required this.messages
  });

  @override
  List<Object?> get props => [messages];
}