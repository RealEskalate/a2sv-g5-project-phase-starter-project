// ignore_for_file: prefer_const_constructors, prefer_const_literals_to_create_immutables

import 'dart:io';
import 'package:ecommerce_app_ca_tdd/extra/icon_img.dart';
import 'package:ecommerce_app_ca_tdd/features/product/domain/entities/product_entity.dart';
import 'package:ecommerce_app_ca_tdd/features/product/presentation/bloc/add/add_event.dart';
import 'package:ecommerce_app_ca_tdd/features/product/presentation/bloc/add/add_state.dart';
import 'package:ecommerce_app_ca_tdd/features/product/presentation/bloc/home_bloc.dart';
import 'package:ecommerce_app_ca_tdd/features/product/presentation/bloc/home_event.dart';
import 'package:ecommerce_app_ca_tdd/features/product/presentation/widgets/bottomnavbar.dart';
import "package:flutter/material.dart";
import 'package:ecommerce_app_ca_tdd/features/product/presentation/pages/details.dart';
import 'package:flutter_bloc/flutter_bloc.dart';
import 'package:get/get.dart';
import 'package:google_fonts/google_fonts.dart';
import 'package:ecommerce_app_ca_tdd/extra/resusetext.dart';
import '../bloc/add/add_bloc.dart';
import 'home.dart';
import 'package:ecommerce_app_ca_tdd/models/product.dart';
import 'package:ecommerce_app_ca_tdd/models/ext_product.dart';
import 'package:image_picker/image_picker.dart';

class AddUpdate extends StatefulWidget {
  const AddUpdate({Key? key}) : super(key: key);

  @override
  _AddUpdateState createState() => _AddUpdateState();
}

class _AddUpdateState extends State<AddUpdate> {
  TextEditingController name_input = TextEditingController();
  TextEditingController category_input = TextEditingController();
  TextEditingController description_input = TextEditingController();
  TextEditingController price_input = TextEditingController();

  File? newImage;
  String pathofimg = '';

  final ImagePicker picker = ImagePicker();

  Future chooseImage() async {
    final pickedImage =
        await picker.pickImage(source: ImageSource.gallery);
    if (pickedImage == null) return;
    setState(() {
      pathofimg = pickedImage.path;
      newImage = File(pickedImage.path);
    });
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      bottomNavigationBar: Container(child: Bottomnavbar()),
      appBar: AppBar(
        automaticallyImplyLeading: false,
        title: Row(
          mainAxisAlignment: MainAxisAlignment.spaceBetween,
          children: [
            IconButton(
                onPressed: () {
                  Navigator.pop(context);
                },
                icon: Icon(
                  Icons.arrow_back_ios_new,
                  color: Color.fromARGB(255, 63, 81, 243),
                  size: 20,
                )),
            Center(
              child: Text("Add Product", style: GoogleFonts.poppins()),
            ),
            SizedBox(
              width: 60,
              height: 60,
            )
          ],
        ),
      ),

      body: SingleChildScrollView(
        padding: EdgeInsets.symmetric(horizontal: 16.0),
        child: Column(
          children: [
            // Upload Image Part Start
            Container(
              margin: EdgeInsets.only(top: 16),
              decoration: BoxDecoration(
                  color: Color.fromRGBO(243, 243, 243, 1),
                  border: Border.all(
                      color: Color.fromRGBO(221, 221, 221, 1), width: 2),
                  borderRadius: BorderRadius.circular(16)),
              child: SizedBox(
                width: double.infinity,
                height: 190,
                child: GestureDetector(
                  onTap: chooseImage,
                  child: Stack(
                    children: [
                      if (newImage != null)
                        Center(
                          child: Image.file(
                            newImage!,
                            fit: BoxFit.cover,
                            width: double.infinity,
                            height: double.infinity,
                          ),
                        ),
                      if (newImage == null)
                        Center(
                          child: Padding(
                            padding: const EdgeInsets.all(16.0),
                            child: Column(
                              mainAxisSize: MainAxisSize.min,
                              mainAxisAlignment: MainAxisAlignment.center,
                              children: [
                                Icon(
                                  Icons.image_outlined,
                                  size: 50,
                                  color: Colors.grey,
                                ),
                                SizedBox(height: 16), // Gap between icon and text
                                Text(
                                  "Upload Image",
                                  style: GoogleFonts.poppins(
                                      fontWeight: FontWeight.w500, fontSize: 14),
                                ),
                              ],
                            ),
                          ),
                        ),
                    ],
                  ),
                ),
              ),
            ), // Upload image part ends here

            SizedBox(
              height: 22,
            ),

            // Input Fields
            BlocConsumer<addBloc, ProductState>(
              listener: (context, state) {
                if (state is ProductLoading) {
                  ScaffoldMessenger.of(context).showSnackBar(
                    SnackBar(
                      backgroundColor: Colors.white,
                      content: Row(
                        children: [
                          Text(
                            'Please Wait...',
                            style: GoogleFonts.poppins(color: Colors.black),
                          ),
                        ],
                      ),
                      duration: Duration(seconds: 5),
                    ),
                  );
                }
                if (state is ProductAddedFailure) {
                  ScaffoldMessenger.of(context).showSnackBar(
                    SnackBar(
                      content: Text('Error: ' + state.error),
                      duration: Duration(seconds: 3),
                    ),
                  );
                }
                if (state is ProductAddedSuccess) {
                  ScaffoldMessenger.of(context).showSnackBar(
                    SnackBar(
                      backgroundColor: Colors.black,
                      content: Row(
                        children: [
                          Icon(Icons.thumb_up_rounded, color: Colors.yellow),
                          SizedBox(width: 10),
                          Text(
                            "Successfully Added!",
                            style: TextStyle(color: Colors.white),
                          ),
                        ],
                      ),
                    ),
                  );
                  context.read<HomeBloc>().add(GetProductsEvent());
                  Navigator.pushNamed(context, '/home');
                }
              },
              builder: (context, state) {
                return Padding(
                  padding: const EdgeInsets.symmetric(vertical: 16.0),
                  child: Column(
                    crossAxisAlignment: CrossAxisAlignment.start,
                    children: [
                      reusableTextpar("Name", FontWeight.w500, 14),
                      SizedBox(height: 8),
                      TextField(
                        controller: name_input,
                        decoration: InputDecoration(
                          border: OutlineInputBorder(
                            borderRadius: BorderRadius.circular(8),
                            borderSide: BorderSide.none,
                          ),
                          filled: true,
                          fillColor: Color.fromRGBO(243, 243, 243, 1),
                        ),
                      ),
                      SizedBox(height: 16),
                      reusableTextpar("Category", FontWeight.w500, 14),
                      SizedBox(height: 8),
                      TextField(
                        controller: category_input,
                        decoration: InputDecoration(
                          border: OutlineInputBorder(
                            borderRadius: BorderRadius.circular(8),
                            borderSide: BorderSide.none,
                          ),
                          filled: true,
                          fillColor: Color.fromRGBO(243, 243, 243, 1),
                        ),
                      ),
                      SizedBox(height: 16),
                      reusableTextpar("Price", FontWeight.w500, 14),
                      SizedBox(height: 8),
                      TextField(
                        controller: price_input,
                        keyboardType: TextInputType.number,
                        decoration: InputDecoration(
                          border: OutlineInputBorder(
                            borderRadius: BorderRadius.circular(8),
                            borderSide: BorderSide.none,
                          ),
                          suffixIcon: Icon(Icons.attach_money),
                          filled: true,
                          fillColor: Color.fromRGBO(243, 243, 243, 1),
                        ),
                      ),
                      SizedBox(height: 16),
                      reusableTextpar("Description", FontWeight.w500, 14),
                      SizedBox(height: 8),
                      TextField(
                        controller: description_input,
                        maxLines: 4,
                        decoration: InputDecoration(
                          border: OutlineInputBorder(
                            borderRadius: BorderRadius.circular(8),
                            borderSide: BorderSide.none,
                          ),
                          filled: true,
                          fillColor: Color.fromRGBO(243, 243, 243, 1),
                        ),
                      ),
                      SizedBox(height: 22),

                      // Buttons
                      Container(
                        margin: EdgeInsets.only(bottom: 22, top: 35),
                        child: Column(
                          children: [
                            SizedBox(
                              width: 366,
                              height: 45,
                              child: ElevatedButton(
                                  style: ElevatedButton.styleFrom(
                                    shape: RoundedRectangleBorder(
                                        borderRadius: BorderRadius.all(
                                            Radius.circular(10))),
                                    side: BorderSide(color: Color(0xff3F51F3)),
                                    foregroundColor: Colors.white,
                                    backgroundColor: Color(0xff3F51F3),
                                  ),
                                  onPressed: () {
                                    final createProduct = ProductEntity(
                                        name: name_input.text,
                                        description: description_input.text,
                                        price: double.tryParse(price_input.text) ?? 0.0,
                                        imagePath: pathofimg);
                                    
                                    var addbloc = BlocProvider.of<addBloc>(context);
                                    addbloc.add(AddProductEvent(product: createProduct));
                                  },
                                  child: Text(
                                    "ADD",
                                    style: GoogleFonts.poppins(
                                        fontSize: 14,
                                        fontWeight: FontWeight.w500),
                                  )),
                            ),
                            SizedBox(
                              height: 16,
                            ),
                            SizedBox(
                              width: 366,
                              height: 45,
                              child: OutlinedButton(
                                  style: OutlinedButton.styleFrom(
                                    shape: RoundedRectangleBorder(
                                        borderRadius: BorderRadius.all(
                                            Radius.circular(10))),
                                    side: BorderSide(color: Colors.red),
                                    foregroundColor: Colors.red,
                                  ),
                                  onPressed: () {
                                    Navigator.pushNamed(context, '/home');
                                  },
                                  child: Text(
                                    "Cancel",
                                    style: GoogleFonts.poppins(
                                        fontSize: 14,
                                        fontWeight: FontWeight.w500),
                                  )),
                            ),
                        SizedBox(height: MediaQuery.of(context).size.height * 0.01,),
                          ],
                        ),
                      ),
                    ],
                  ),
                );
              },
            )
          ],
        ),
      ),
    );
  }
}
