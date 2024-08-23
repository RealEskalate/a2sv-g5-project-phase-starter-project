// import 'dart:collection';
// import 'dart:js_interop';

import 'package:ecommerce/features/auth/data/model/user_model.dart';
import 'package:flutter/material.dart';
import 'package:flutter_bloc/flutter_bloc.dart';
import 'package:image_picker/image_picker.dart';
import 'dart:io';

import '../../../../service_locator.dart';
import '../../data/model/product_model.dart';
import '../../domain/usecases/addproduct.dart';
import '../../domain/usecases/getallproduct.dart';
import '../bloc/getallproductbloc/bloc/product_bloc.dart';
import 'homepage.dart';

class AddUpdate extends StatefulWidget {
  final UserModel user;
  const AddUpdate({Key? key, required this.user}) : super(key: key);
  @override
  _AddUpdateState createState() => _AddUpdateState();
}

class _AddUpdateState extends State<AddUpdate> {
  File? _image;
  
  TextEditingController nameController = TextEditingController();
  TextEditingController descriptionController = TextEditingController();
  TextEditingController categoryController = TextEditingController();
  TextEditingController priceController = TextEditingController();
  var productBloc = getIt<ProductBloc>();


  Future<void> _pickImage() async {
    final ImagePicker _picker = ImagePicker();
    final XFile? image = await _picker.pickImage(source: ImageSource.gallery);

    if (image != null) {
      setState(() {
        _image = File(image.path);
      });
    }
  }

  
  var getAllProductUsecase = getIt<GetAllProductUsecase>();
  var addProductUsecase = getIt<AddProductUsecase>();

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      backgroundColor: Colors.white,
      body: BlocProvider(
        create: (context) => getIt<ProductBloc>(),
        child: SingleChildScrollView(
          child: Padding(
            padding: EdgeInsets.only(left: 20.0, top: 30, right: 20.0),
            child: Column(
              crossAxisAlignment: CrossAxisAlignment.start,
              children: [
                BlocListener<ProductBloc, ProductState>(
                  bloc: productBloc,
                  listener: (context, state) {
                    if (state is added) {
                      ScaffoldMessenger.of(context).showSnackBar(
                        SnackBar(
                          content: Text('Product Added Successfully'),
                        ),
                      );
                      Navigator.push(
                        context,
                        MaterialPageRoute(
                          builder: (context) => MyHomePage(title: '', user: widget.user),
                        ),
                      );
                    }
                     else if (state is addfailure){
                      ScaffoldMessenger.of(context).showSnackBar(
                        SnackBar(
                          content: Text('Product Added UnSuccessfully'),
                        ),
                      );
                      } 
                    
                  },
                  child: Row(
                    children: [
                      IconButton(
                        icon:
                            Icon(Icons.arrow_back_ios, color: Color(0xFF3f51f3)),
                        onPressed: () {
                          Navigator.pop(context);
                        },
                      ),
                      const SizedBox(width: 80),
                      const Text(
                        'Add Products',
                        style: TextStyle(
                          fontSize: 20,
                          // fontWeight: FontWeight.bold,
                        ),
                      ),
                    ],
                  ),
                ),
                SizedBox(height: 20),
                // Container for image upload
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
                    staticText: '',
                    maxLines: 1,
                    flag: false,
                    controller: nameController),

                const SizedBox(
                  height: 10,
                ),
                const Text('price',
                    style: TextStyle(
                      fontSize: 18,
                    )),
                CustomTextInputField(
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

                CustomTextInputField(
                  staticText: '',
                  maxLines: 4,
                  flag: true,
                  controller: descriptionController,
                ),
                SizedBox(
                  height: 10,
                ),
                SizedBox(
                  width: double.infinity, // Full width of the parent
                  child: OutlinedButton(
                    style: OutlinedButton.styleFrom(
                      backgroundColor: Color(0xFF3f51f3), // Button color
                      minimumSize: const Size(
                          double.infinity, 50), // Minimum width and height
                      shape: RoundedRectangleBorder(
                        borderRadius:
                            BorderRadius.circular(8), // Border radius
                      ),
                    ),
                    onPressed: () {
                      productBloc.add(AddProductEvent(
                          product: ProductModel(
                        id: "123",
                        name: nameController.text,
                        // category: categoryController.text,
                        price: double.parse(priceController.text),
                        description: descriptionController.text,
                        image: _image!.path,
                        seller: widget.user,
                      )));
                    },
                    child: const Text(
                      'ADD',
                      style: TextStyle(
                        color: Colors.white,
                      ),
                    ),
                  ),
                ),
                SizedBox(
                  height: 10,
                ),
                SizedBox(
                  width: double.infinity, // Full width of the parent
                  child: OutlinedButton(
                    style: OutlinedButton.styleFrom(
                      foregroundColor: Colors.red,
                      side: BorderSide(color: Colors.red), // Button color
                      minimumSize: const Size(
                          double.infinity, 50), // Minimum width and height
                      shape: RoundedRectangleBorder(
                        borderRadius: BorderRadius.circular(8), // Border radius
                      ),
                    ),
                    onPressed: () {},
                    child: const Text(
                      'DELETE',
                      style: TextStyle(
                        color: Colors.red,
                      ),
                    ),
                  ),
                ),
              ],
            ),
          ),
        ),
      ),
    );
  }
}

class CustomTextInputField extends StatelessWidget {
  final String? labelText;
  final int? maxLines;
  final String? staticText;
  final bool? flag;
  TextEditingController controller = TextEditingController();

  CustomTextInputField(
      {this.labelText,
      this.maxLines = 1,
      required this.staticText,
      required this.flag,
      required this.controller});

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
              labelText: labelText,
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
