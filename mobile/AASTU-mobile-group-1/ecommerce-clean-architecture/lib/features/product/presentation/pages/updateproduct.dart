import 'package:ecommerce/features/auth/data/model/user_model.dart';
import 'package:flutter/material.dart';
import 'package:flutter_bloc/flutter_bloc.dart';

import 'package:image_picker/image_picker.dart';
import 'dart:io';

import '../../../../service_locator.dart';
import '../../data/model/product_model.dart';
import '../../domain/entities/product.dart';
import '../bloc/getallproductbloc/bloc/product_bloc.dart';
import 'homepage.dart';

class Updateproduct extends StatefulWidget {
  final UserModel user;
  Updateproduct({super.key, required this.sampleproduct, required this.user});
  Productentity sampleproduct;
  @override
  _Updateproductstate createState() => _Updateproductstate();
}

class _Updateproductstate extends State<Updateproduct> {
  File? _image;
  TextEditingController nameController = TextEditingController();
  TextEditingController priceController = TextEditingController();
  TextEditingController descriptionController = TextEditingController();
  var productbloc = getIt<ProductBloc>();

  @override
  void initState() {
    super.initState();
    nameController.text = widget.sampleproduct.name;
    priceController.text = widget.sampleproduct.price.toString();
    descriptionController.text = widget.sampleproduct.description;
  }

  Future<void> _pickImage() async {
    final ImagePicker _picker = ImagePicker();
    final XFile? image = await _picker.pickImage(source: ImageSource.gallery);

    if (image != null) {
      setState(() {
        _image = File(image.path);
      });
    }
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      backgroundColor: Colors.white,
      body: BlocConsumer<ProductBloc, ProductState>(
        listener: (context, state) {
              if (state is updated) {
                Navigator.push(
                    context,
                    MaterialPageRoute(
                        builder: (context) => MyHomePage(
                              title: '', user: widget.user,
                            )));
                ScaffoldMessenger.of(context).showSnackBar(
                  SnackBar(
                    content: Text('Product Updated'),
                    // duration: Duration(seconds: 2),
                  ),
                );
              } else if (state is updatefailure) {
                ScaffoldMessenger.of(context).showSnackBar(
                  SnackBar(
                    content: Text(state.message),
                    // duration: Duration(seconds: 2),
                  ),
                );
              }
            },
      
        builder: (context, state) {
          return  Expanded(
              child: SingleChildScrollView(
                child: Padding(
                  padding: EdgeInsets.only(left: 20.0, top: 30, right: 20.0),
                  child: Column(
                    crossAxisAlignment: CrossAxisAlignment.start,
                    children: [
                      Row(
                        children: [
                          IconButton(
                            icon: Icon(Icons.arrow_back_ios,
                                color: Color(0xFF3f51f3)),
                            onPressed: () {
                              Navigator.pop(context);
                            },
                          ),
                          const SizedBox(width: 70),
                          const Text(
                            'Update Product',
                            style: TextStyle(
                              fontSize: 20,
                              // fontWeight: FontWeight.bold,
                            ),
                          ),
                        ],
                      ),
                      SizedBox(height: 20),
                      Container(
                        height: 250,
                        decoration: BoxDecoration(
                          color: Colors.grey[200],
                          borderRadius: BorderRadius.circular(12.0),
                          boxShadow: const [
                            BoxShadow(
                              color: Colors.black12,
                              blurRadius: 8.0,
                              offset: Offset(0, 4),
                            ),
                          ],
                        ),
                        child: Column(
                          mainAxisAlignment: MainAxisAlignment.center,
                          children: [
                            GestureDetector(
                              onTap: _pickImage,
                              child: Icon(
                                Icons.image_search_outlined,
                                color: Colors.grey[600],
                                size: 50,
                              ),
                            ),
                            _image == null
                                ? Center(
                                    child: Text(
                                      'Upload Image',
                                      style: TextStyle(color: Colors.grey[600]),
                                    ),
                                  )
                                : Image.file(
                                    _image!,
                                    height: 200,
                                    width: double.infinity,
                                    fit: BoxFit.cover,
                                  ),
                          ],
                        ),
                      ),
                      const SizedBox(
                        height: 10,
                      ),
                      const Text(
                        'name',
                        style: TextStyle(
                          fontSize: 18,
                        ),
                      ),
                      CustomTextInputField(
                        initialvalue: widget.sampleproduct.name,
                        staticText: '',
                        maxLines: 1,
                        flag: false,
                        controller: nameController,
                      ),
                      const SizedBox(
                        height: 10,
                      ),
                      const Text('price',
                          style: TextStyle(
                            fontSize: 18,
                          )),
                      CustomTextInputField(
                        initialvalue: widget.sampleproduct.price.toString(),
                        staticText: '\$',
                        maxLines: 1,
                        flag: false,
                        controller: priceController,
                      ),
                      const SizedBox(
                        height: 10,
                      ),
                      const Text('description',
                          style: TextStyle(
                            fontSize: 18,
                          )),
                      const SizedBox(
                        height: 10,
                      ),
                      CustomTextInputField(
                        initialvalue: widget.sampleproduct.description,
                        staticText: '',
                        maxLines: 5,
                        flag: true,
                        controller: descriptionController,
                      ),
                      SizedBox(
                        height: 40,
                      ),
                      SizedBox(
                        width: double.infinity, // Full width of the parent
                        child: OutlinedButton(
                          style: OutlinedButton.styleFrom(
                            backgroundColor: Color(0xFF3f51f3), // Button color
                            minimumSize: const Size(double.infinity,
                                50), // Minimum width and height
                            shape: RoundedRectangleBorder(
                              borderRadius:
                                  BorderRadius.circular(8), // Border radius
                            ),
                          ),
                          onPressed: () {
                            productbloc.add(UpdateProductEvent(
                                product: ProductModel(
                              id: widget.sampleproduct.id,
                              name: nameController.text,
                              price: double.parse(priceController.text),
                              description: descriptionController.text,
                              image: _image.toString(),
                              seller: widget.user,
                            )));

                            // Navigator.push(context, MaterialPageRoute(builder: (context)=>MyHomePage(title: '',)));
                          },
                          child: const Text(
                            'UPDATE',
                            style: TextStyle(
                              color: Colors.white,
                            ),
                          ),
                        ),
                      ),
                    ],
                  ),
                ),
              ),
            );
        },
      ),
    );
  }
}

class CustomTextInputField extends StatelessWidget {
  final int? maxLines;
  final String? staticText;
  final bool? flag;
  final String? initialvalue;
  final TextEditingController controller;

  CustomTextInputField(
      {this.maxLines = 1,
      required this.staticText,
      required this.flag,
      required this.controller,
      required this.initialvalue});

  @override
  Widget build(BuildContext context) {
    return Container(
      height: flag == true ? 110 : 35,
      child: Stack(
        children: [
          TextField(
            maxLines: maxLines,
            controller: controller,
            decoration: InputDecoration(
              labelText: '',
              border: OutlineInputBorder(
                borderRadius: BorderRadius.circular(5.0),
                borderSide: BorderSide.none,
              ),
              filled: true,
              fillColor: Color(0xFFF3F3F3),
            ),
          ),
          if (staticText != null)
            Positioned(
                right: 10,
                child: Text(
                  staticText!,
                  style: TextStyle(
                    color: Colors.grey,
                    fontSize: 20,
                  ),
                ))
        ],
      ),
    );
  }
}


// class CustomTextInputField extends StatefulWidget {
//   final String? labelText;
//   final int? maxLines;
//   final String? staticText;
//   final bool? flag;
//   final TextEditingController controller;

//   CustomTextInputField({
//     required this.labelText,
//     this.maxLines = 1,
//     required this.staticText,
//     required this.flag,
//     required this.controller,
//   });

//   @override
//   _CustomTextInputFieldState createState() => _CustomTextInputFieldState();
// }

// class _CustomTextInputFieldState extends State<CustomTextInputField> {
//   var controller = widget.controller;

//   @override
//   void initState() {
//     super.initState();
//     controller = TextEditingController(text: widget.labelText);
//   }

//   @override
//   void dispose() {
//     controller.dispose();
//     super.dispose();
//   }

//   @override
//   Widget build(BuildContext context) {
//     return Container(
//       height: widget.flag == true ? 110 : 35,
//       child: Stack(
//         children: [
//           TextField(
//             maxLines: widget.maxLines,
//             controller : controller,
//             decoration: InputDecoration(
//               // labelText: widget.labelText,
//               border: OutlineInputBorder(
//                 borderRadius: BorderRadius.circular(5.0),
//                 borderSide: BorderSide.none,
//               ),
//               filled: true,
//               fillColor: Color(0xFFF3F3F3),
//             ),
          
//           ),
//           if (widget.staticText != null)
//             Positioned(
//               right: 10,
//               top: 12, // Adjust the top position to align the text correctly
//               child: Text(
//                 widget.staticText!,
//                 style: TextStyle(
//                   color: Colors.grey,
//                   fontSize: 20,
//                 ),
//               ),
//             ),
//         ],
//       ),
//     );
//   }
// }
