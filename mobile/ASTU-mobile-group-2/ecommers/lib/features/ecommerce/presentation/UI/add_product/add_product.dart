import 'dart:io';

import 'package:flutter/material.dart';
import 'package:flutter_bloc/flutter_bloc.dart';
import 'package:flutter_easyloading/flutter_easyloading.dart';

import '../../../../../core/const/width_height.dart';
import '../../state/image_input_display/image_bloc.dart';
import '../../state/image_input_display/image_event.dart';
import '../../state/image_input_display/image_state.dart';
import '../../state/input_button_activation/bottum_state.dart';
import '../../state/input_button_activation/button_bloc.dart';
import '../../state/input_button_activation/button_event.dart';
import '../../state/product_bloc/product_bloc.dart';
import '../../state/product_bloc/product_event.dart';
import '../../state/product_bloc/product_state.dart';
import 'add_delete_button.dart';
import 'convert_input_to_json.dart';
import 'image_field.dart';
import 'input_border.dart';

class AddProduct extends StatefulWidget {

  const AddProduct({
    super.key,
 
    
  });

  @override
  State<AddProduct> createState() => _AddProductState();
}

class _AddProductState extends State<AddProduct> {
  @override
  Widget build(BuildContext context) {
    Map<String, dynamic> data =
        ModalRoute.of(context)!.settings.arguments as Map<String, dynamic>;
      // data = data == {}?widget.datas:data;
      
    double width = WidthHeight.screenWidth(context);
    double height = WidthHeight.screenHeight(context);
    File? localImage;
    return SafeArea(
      child: Scaffold(
        appBar: AppBar(
          backgroundColor: Colors.white,
          title:  data['type'] != 1?const Text('Add Product'):const Text('Update Product'),
          centerTitle: true,
          leading: GestureDetector(
            key: const Key('back from add page'),
            onTap: () => {
              setState(() {
                localImage = null;
                EasyLoading.dismiss();
                context.read<ButtonBloc>().add(InsertInput(type: 0,id:data['id']));
                Navigator.pop(context);
              })
            },
            child: BlocListener<ProductBloc, ProductState>(
              listener: (context, state) {
                if (state is ProductErrorState) {
                  ScaffoldMessenger.of(context).showSnackBar(
                    const SnackBar(
                      content: Text('try again'),
                    ),
                  );
                  EasyLoading.showSuccess('try again');
                  EasyLoading.dismiss();
                } else if (state is SuccessDelete) {
                  ScaffoldMessenger.of(context).showSnackBar(
                    const SnackBar(
                      content: Text('success'),
                    ),
                  );
                  context.read<ProductBloc>().add(const LoadAllProductEvent());
                  EasyLoading.showSuccess('success');
                  EasyLoading.dismiss();
                  Navigator.popUntil(context, ModalRoute.withName('/home'));
                }
              },
              child: Container(
                width: 10,
                height: 10,
                margin: const EdgeInsets.only(left: 10),
                decoration: BoxDecoration(
                    borderRadius: BorderRadius.circular(30), color: Colors.white),
                child: const Icon(
                  Icons.arrow_back_ios_new,
                  color: Colors.blue,
                ),
              ),
            ),
          ),
        ),
        body: SingleChildScrollView(
          child: BlocListener<ButtonBloc, BottumState>(
            listener: (context, state) {
              if (state is AddErrorState) {
                EasyLoading.showError('Network error');
                
                ScaffoldMessenger.of(context).showSnackBar(
                  const SnackBar(
                    content: Text('try again'),
                  ),
                );
      
                EasyLoading.dismiss();
                context.read<ButtonBloc>().add(InsertInput(type: data['type'],id: data['id']));
              } else if (state is SuccessAddProduct) {
                if (state.add) {
                  context.read<ProductBloc>().add(const LoadAllProductEvent());
                  EasyLoading.showSuccess('Great Success!');
                  ScaffoldMessenger.of(context).showSnackBar(
                    const SnackBar(
                      content: Text('success'),
                    ),
                  );
                  EasyLoading.dismiss();
                  Navigator.popUntil(context, (Route<dynamic> route) {
                    return route.settings.name == '/home';
                  });
                } else {
                  EasyLoading.dismiss();
                  ScaffoldMessenger.of(context).showSnackBar(
                    const SnackBar(
                      content: Text('try again'),
                    ),
                  );
                }
              }
            },
            child: Container(
              width: double.infinity,
              color: Colors.white,
              padding: const EdgeInsets.fromLTRB(20, 15, 20, 10),
              child: Column(
                crossAxisAlignment: CrossAxisAlignment.start,
                children: [
                  BlocBuilder<ImageBloc, ImageState>(
                    builder: (context, state) {
                      if (state is OnImageSelect) {
                        localImage = state.file;
                        context
                            .read<ButtonBloc>()
                            .add(InsertInput(image: state.file, tag: 'image',type: data['type'],id:data['id']));
                        state.file = null;
                        state = InputIntialState();
                      }
                      
                      return GestureDetector(
                         key: const Key('image selector'),
                          onTap: () => {
                                context.read<ImageBloc>().add(SelectImageEvent()),
                              },
                          child: ImageField(
                            hight: (0.3*height).toInt(),
                            check: true,
                            text: 'Photo',
                            width: (width).toInt(),
                            imageUrl: data['imageUrl'],
                            localImage: localImage,
                          ));
                    },
                  ),
                  const SizedBox(
                    height: 10,
                  ),
                  
                  const Text('name'),
                  const SizedBox(
                    height: 5,
                  ),
                  IinputBorder(
                      hight: (0.069*height).toInt(),
                      check: false,
                      text: 'name',
                      width: (width).toInt(),
                      placeHolder: data['name'],
                      data: data,
                      ),
                  const SizedBox(
                    height: 10,
                  ),
                  const Text('catagory'),
                  const SizedBox(
                    height: 5,
                  ),
                  IinputBorder(
                      hight: (0.069*height).toInt(),
                      check: false,
                      text: 'catagory',
                      width: (width).toInt(),
                      placeHolder: data['name'],data: data,),
                  const SizedBox(
                    height: 10,
                  ),
                  const Text('price'),
                  const SizedBox(
                    height: 5,
                  ),
                  IinputBorder(
                      hight: 56,
                      check: false,
                      text: 'price',
                      width: (width).toInt(),
                      placeHolder: data['price'] > 0?data['price'].toString():'',data: data,),
                  const SizedBox(
                    height: 10,
                  ),
                  const Text('description'),
                  const SizedBox(
                    height: 5,
                  ),
                  IinputBorder(
                      hight: 140,
                      check: false,
                      text: 'description',
                      width: (width).toInt(),
                      placeHolder: data['disc'],data: data,),
                  const SizedBox(
                    height: 15,
                  ),
                  ConvertInputToJson(
                    data: data,
                  ),
                  const SizedBox(
                    height: 10,
                  ),
                  !(data['id'].isEmpty)
                      ? GestureDetector(
                          onTap: () => {
                            EasyLoading.showProgress(0.3, status: 'Deleting...'),
                            context
                                .read<ProductBloc>()
                                .add(DeleteProductEvent(id: data['id'])),
                          },
                          child: const AddDeleteButton(
                            color: Colors.white,
                            text: 'DELETE',
                            borderColor: Colors.red,
                          ),
                        )
                      : const Text(''),
                ],
              ),
            ),
          ),
        ),
      ),
    );
  }
}
