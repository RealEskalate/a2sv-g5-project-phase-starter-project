


import 'package:flutter_bloc/flutter_bloc.dart';

import '../../../Domain/usecase/ecommerce_usecase.dart';
import 'image_event.dart';
import 'image_state.dart';

class ImageBloc extends Bloc<ImageEvent,ImageState>{
  final EcommerceUsecase ecommerceUsecase;
  ImageBloc ({
    required this.ecommerceUsecase
  }):super(InputIntialState()){
    // on<On

  on<SelectImageEvent>(
      (event,emit) async {
        emit(ImageLoadingState());
        final result = await ecommerceUsecase.selectImage();
        result.fold(
          (failure){
            emit(ErrorState(messages: 'try again'));
          }, 
          (data) {
            emit(OnImageSelect(image: data['image'],file: data['file']));
          }
        );
      }
    );
}
}