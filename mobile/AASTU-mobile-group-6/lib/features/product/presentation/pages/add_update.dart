// ignore_for_file: prefer_const_constructors, prefer_const_literals_to_create_immutables

import 'dart:math';
import 'dart:io';
import 'package:ecommerce_app_ca_tdd/extra/icon_img.dart';
import 'package:ecommerce_app_ca_tdd/features/product/domain/entities/product_entity.dart';
import 'package:ecommerce_app_ca_tdd/features/product/presentation/bloc/add/add_event.dart';
import 'package:ecommerce_app_ca_tdd/features/product/presentation/bloc/add/add_state.dart';
import 'package:ecommerce_app_ca_tdd/features/product/presentation/bloc/home_bloc.dart';
import 'package:ecommerce_app_ca_tdd/features/product/presentation/bloc/home_event.dart';
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

  //
  final ImagePicker picker = ImagePicker();

  Future chooseImage() async {
    final pickedImage =
        await ImagePicker().pickImage(source: ImageSource.gallery);
    if (pickedImage == null) return;
    setState(() {
      pathofimg = pickedImage.path;
      newImage = File(pickedImage.path);
    });
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
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
            const Center(
              child: Text("Add Product"),
            ),
            const SizedBox(
              height: 60,
              width: 60,
            )
          ],
        ),
      ),

      body: SingleChildScrollView(
        child: Column(
          children: [
            // Upload Image Part Start
            Container(
              margin: EdgeInsets.only(left: 32, right: 32, top: 23),
              decoration: BoxDecoration(
                  color: Color.fromRGBO(243, 243, 243, 1),
                  border: Border.all(
                      color: Color.fromRGBO(221, 221, 221, 1), width: 2),
                  borderRadius: BorderRadius.circular(16)),
              child: SizedBox(
                width: 360,
                height: 190,
                child: GestureDetector(
                  onTap: chooseImage,
                  child: Column(
                    mainAxisAlignment: MainAxisAlignment.center,
                    children: [
                      if (newImage == null)
                        IconButton(
                            onPressed: () {},
                            icon: newImage == null
                                ? Icon(
                                    Icons.image_outlined,
                                    size: 50,
                                  )
                                : ImagePickerIconButton()),
                      SizedBox(
                        height: 30,
                      ),
                      Center(
                          child: reusableTextpar(
                              "upload image", FontWeight.w500, 14)),
                    ],
                  ),
                ),
              ),
            ), // Upload image part ends here

            SizedBox(
              height: 22,
            ),

            // Name Entry Field Start
            BlocConsumer<addBloc, ProductState>(
              listener: (context, state) {
                  if (state is ProductLoading){
                      ScaffoldMessenger.of(context).showSnackBar(
                        SnackBar(
                          backgroundColor: Colors.white,
                          content: Row(
                            children: [
                              Text('Please Wait...',style: GoogleFonts.poppins(color: Colors.black),),
                            ],
                          ),
                          duration: Duration(seconds: 5),
                        ),
                      );
            
                  }
                  if (state is ProductAddedFailure){
                              
                        ScaffoldMessenger.of(context).showSnackBar(
                          SnackBar(
                            content: Text('Error:' + state.error),
                            duration: Duration(seconds: 3),
                          ),
                        );
                      
                  }
                  if(state is ProductAddedSuccess){
                    
                    ScaffoldMessenger.of(context).showSnackBar(
                    SnackBar(
                      backgroundColor: Colors.black,
                  content: Row(
                    // 
                    children: [
                      Icon(Icons.thumb_up_rounded,color: Colors.yellow,),
                      SizedBox(width: 10,),
                      Text("Succesfully Added !!!!",style: TextStyle(color: Colors.white),),
                    ],
                  ),
                      
                    ));
                    context.read<HomeBloc>().add(GetProductsEvent());

                    Navigator.pushNamed(context,'/home');
                  }
                
                                  
                                    
              },
              builder: (context, state) {
                return SingleChildScrollView(
                  child: Container(
                      child: Column(
                    mainAxisAlignment: MainAxisAlignment.start,
                    crossAxisAlignment: CrossAxisAlignment.start,
                    children: [
                      reusableTextpar("name", FontWeight.w500, 14),
                      SizedBox(
                        height: 8,
                      ),
                      SizedBox(
                          height: 50,
                          width: 360,
                          child: TextField(
                            controller: name_input,
                            maxLines: 4,
                            decoration: InputDecoration(
                                border: InputBorder.none,
                                filled: true,
                                fillColor: Color.fromRGBO(243, 243, 243, 1)),
                          )),
                      reusableTextpar("category", FontWeight.w500, 14),
                      SizedBox(
                        height: 8,
                      ),
                      SizedBox(
                          height: 50,
                          width: 360,
                          child: TextField(
                            controller: category_input,
                            decoration: InputDecoration(
                                border: InputBorder.none,
                                filled: true,
                                fillColor: Color.fromRGBO(243, 243, 243, 1)),
                          )),
                      reusableTextpar("price", FontWeight.w500, 14),
                      SizedBox(
                        height: 8,
                      ),
                      SizedBox(
                          height: 50,
                          width: 360,
                          child: TextField(
                            controller: price_input,
                            decoration: InputDecoration(
                              suffixIcon: Icon(Icons.attach_money),
                                border: InputBorder.none,
                                hintText: "",
                                filled: true,
                                fillColor: Color.fromRGBO(243, 243, 243, 1)),
                          )),
                      reusableTextpar("description", FontWeight.w500, 14),
                      SizedBox(
                        height: 8,
                      ),
                      SizedBox(
                          height: 150,
                          width: 360,
                          child: TextField(
                            controller: description_input,
                            maxLines: 4,
                            decoration: InputDecoration(
                                border: InputBorder.none,
                                filled: true,
                                fillColor: Color.fromRGBO(243, 243, 243, 1)),
                          )),
                      // End of Input Fields

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
                                    // overlayColor: Colors.red,
                                    foregroundColor: Colors.white,
                                    backgroundColor: Color(0xff3F51F3),
                                  ),
                                  onPressed: () {
                                    final createProduct = ProductEntity(
                                        name: name_input.text,
                                        description: description_input.text,
                                        price: double.parse(price_input.text),
                                        imagePath: pathofimg);
                                    
                                    var addbloc = BlocProvider.of<addBloc>(context);
                                    addbloc.add(AddProductEvent(
                                        product: createProduct));
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
                                    overlayColor: Colors.red,
                                    foregroundColor: Colors.red,
                                  ),
                                  onPressed: () {
                                    Navigator.pushNamed(context, '/home');
                                  },
                                  child: Text("Cancel",
                                      style: GoogleFonts.poppins(
                                          fontSize: 14,
                                          fontWeight: FontWeight.w500))),
                            ),
                          ],
                        ),
                      ),
                    ],
                  )),
                );
              },
            )
          ],
        ),
      ),
      // body: Text()
    );
  }
} // End of AddUpdate Class
