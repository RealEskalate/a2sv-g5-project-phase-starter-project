import 'dart:io';

// ignore: depend_on_referenced_packages
import 'package:bloc/bloc.dart';
import 'package:dartz/dartz.dart';
import 'package:equatable/equatable.dart';

import '../../../../../core/errors/failures/failure.dart';
import '../../../../../core/validator/validator.dart';

part 'input_validation_state.dart';

class InputValidationCubit extends Cubit<InputValidationState> {
  final InputDataValidator inputDataValidator;
  InputValidationCubit(this.inputDataValidator)
      : super(InputValidationInitial());

  void checkChanges(List<dynamic> typeAndVal) {
    Map<String, Either<Failure, bool>> correspond = {
      'Name': inputDataValidator.checkNameOrCatagory(typeAndVal[1]),
      'Catagory': inputDataValidator.checkNameOrCatagory(typeAndVal[1]),
      'Price': inputDataValidator.checkPrice(typeAndVal[1])
    };

    if (!correspond.containsKey(typeAndVal[0])) return;

    final myCheck = correspond[typeAndVal[0]];

    List<dynamic> nameData = [state.name, state.nameMessage];
    List<dynamic> catagoryData = [state.catagory, state.catagoryMessage];
    List<dynamic> priceData = [state.price, state.priceMessage];
    File? imageUrl = state.imageUrl;
    Map<String, List<dynamic>> correspondList = {
      'Name': nameData,
      'Catagory': catagoryData,
      'Price': priceData
    };
    myCheck?.fold((failure) {
      correspondList[typeAndVal[0]]![0] = false;
      correspondList[typeAndVal[0]]![1] = failure.message;
    }, (data) {
      correspondList[typeAndVal[0]]![0] = true;
    });

    emit(InputValidatedState(
      name: nameData[0],
      catagory: catagoryData[0],
      price: priceData[0],
      imageUrl: imageUrl,
      nameMessage: nameData[1],
      catagoryMessage: catagoryData[1],
      priceMessage: priceData[1],
    ));
  }

  void setImage(File urls) {
    List<dynamic> nameData = [state.name, state.nameMessage];
    List<dynamic> catagoryData = [state.catagory, state.catagoryMessage];
    List<dynamic> priceData = [state.price, state.priceMessage];

    emit(InputValidatedState(
      name: nameData[0],
      catagory: catagoryData[0],
      price: priceData[0],
      imageUrl: urls,
      nameMessage: nameData[1],
      catagoryMessage: catagoryData[1],
      priceMessage: priceData[1],
    ));
  }

  void refresh() {
    emit(InputValidationInitial());
  }
}
