
import 'dart:io';

import 'package:flutter_bloc/flutter_bloc.dart';

import '../../../../../core/utility/input_converter.dart';
import '../../../Domain/usecase/ecommerce_usecase.dart';
import 'bottum_state.dart';
import 'button_event.dart';

class ButtonBloc  extends Bloc<ButtonEvent,BottumState>{
  final EcommerceUsecase ecommerceUsecase;
  ButtonBloc ({required this.ecommerceUsecase}):super(IntialState()){
      String name = '';
      String price = '';
      String description = '';
      String id = '';
     
      File? file;
      on<InsertInput>(
      (event,emit) async {
        if (event.tag == 'name'){
          name = event.name;
        } else if (event.tag == 'price'){
          price = event.price;
        } else if (event.tag == 'description'){
          description = event.description;
        }  else if (event.tag == 'image'){
          file = event.image;
        } else{
          name = name.isEmpty? event.name: name;
          price = price.isEmpty? event.price: price;
          description = description.isEmpty? event.description: description;
          id = event.id;
          file = event.image?? file;
        }
       
        if (checkInput(name, price, description, file,event.type)){
           emit(OnButtonActivate(isActivate: true));
        }
        else if (checkInput(name, price, description, file,event.type)){
          emit(OnButtonActivate(isActivate: true));
        } else {
          emit(OnButtonActivate(isActivate: false));
        }

      }
    );

    on<AddProductEvent>(
      (event,emit) async {
        price = price.isEmpty? '0.0':price;
     final Map<String, dynamic> data = {

          'name': name,
          'price': price,
          'description': description,
          'file': file
        };

      // Convert the Map to a JSON string using jsonEncode
      // final String jsonData = jsonEncode(data);
      emit(AddLoadingState());
      final result = await ecommerceUsecase.addProducts(data);
      result.fold(
        (failure){
          emit(AddErrorState(messages: 'try again'));
        }, 
        (data) {
         
          emit(SuccessAddProduct(add: data));
        }
      );
      }
    );
    on<UpdateProductEvent>(
      (event,emit) async {
        price = price.isEmpty? '0.0':price;
        final double priceDouble = double.parse(price);
        final Map<String, dynamic> data = {

          'name': name,
          'price': priceDouble,
          'description': description,
        };
        emit(AddLoadingState());
        final result = await ecommerceUsecase.editProduct(id,data);
        result.fold(
          (failure){
            emit(AddErrorState(messages: 'try again'));
          }, 
          (data) {
            emit(SuccessAddProduct(add: data));
          }
        );
      }
    );
  }
}



