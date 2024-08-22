import 'dart:io';

import 'package:ecommerce_app_ca_tdd/features/product/data/models/product_models.dart';
import 'package:ecommerce_app_ca_tdd/features/product/presentation/bloc/home_bloc.dart';
import 'package:ecommerce_app_ca_tdd/features/product/presentation/bloc/home_event.dart';
import 'package:ecommerce_app_ca_tdd/features/product/presentation/bloc/update/bloc/update_bloc.dart';
import 'package:ecommerce_app_ca_tdd/features/product/presentation/bloc/update/bloc/update_event.dart';
import 'package:ecommerce_app_ca_tdd/features/product/presentation/bloc/update/bloc/update_state.dart';
import 'package:flutter/material.dart';
import 'package:ecommerce_app_ca_tdd/features/product/presentation/pages/home.dart';
import 'package:flutter_bloc/flutter_bloc.dart';
import 'package:get/get.dart';
import 'package:google_fonts/google_fonts.dart';
import 'package:image_picker/image_picker.dart';
import 'package:ecommerce_app_ca_tdd/extra/resusetext.dart';

class UpdatePage extends StatefulWidget {
  final ProductModel product;
  const UpdatePage({super.key, required this.product});

  @override
  State<UpdatePage> createState() => _UpdatePageState();
}

class _UpdatePageState extends State<UpdatePage> {
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
  bool _isSubmitting = false;

  @override
  void initState() {
    super.initState();
    name_input = TextEditingController(text: widget.product.name);
    description_input = TextEditingController(text: widget.product.description);
    price_input = TextEditingController(text: widget.product.price.toString());
  }

  void _addHandler(BuildContext context) {
    setState(() {
      _isSubmitting = true;
    });}

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
              child: Text("Update Product"),
            ),
            const SizedBox(
              height: 60,
              width: 60,
            )
          ],
        ),
      ),

      body: BlocConsumer<UpdateBloc, UpdateState>(
        listener: (context, state) {
          if (state is UpdateSuccess) {
            ScaffoldMessenger.of(context).showSnackBar(SnackBar(
                content: Text('Success: Product Updated Successfully')));
            context.read<HomeBloc>().add(GetProductsEvent());
            Navigator.pushNamed(context, '/home');
          } else if (state is UpdateFailiure) {
            ScaffoldMessenger.of(context)
                .showSnackBar(SnackBar(content: Text(state.error)));
          }
        },
        builder: (context, state) {
          return SingleChildScrollView(
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
                  child: GestureDetector(
                    onTap: chooseImage,
                    child: SizedBox(
                      width: 360,
                      height: 190,
                      child: Image(
                        image: NetworkImage(widget.product.imagePath),
                        fit: BoxFit.cover,
                      ),
                    ),
                  ),
                ), // Upload image part ends here

                SizedBox(
                  height: 22,
                ),

                // Name Entry Field Start
                SingleChildScrollView(
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
                              child: BlocBuilder<UpdateBloc, UpdateState>(
                                builder: (context, state) {
                                  return ElevatedButton(
                                      style: ElevatedButton.styleFrom(
                                        shape: RoundedRectangleBorder(
                                            borderRadius: BorderRadius.all(
                                                Radius.circular(10))),
                                        side: BorderSide(
                                            color: Color(0xff3F51F3)),
                                        // overlayColor: Colors.red,
                                        foregroundColor: Colors.white,
                                        backgroundColor: Color(0xff3F51F3),
                                      ),
                                      child: Text(
                                        "Update",
                                        style: GoogleFonts.poppins(
                                            fontSize: 14,
                                            fontWeight: FontWeight.w500),
                                      ),
                                      onPressed: () {
                                        final updateModel = ProductModel(
                                            id: widget.product.id,
                                            name: name_input.text,
                                            description: description_input.text,
                                            price:
                                                double.parse(price_input.text),
                                            imagePath:
                                                widget.product.imagePath);
                                                if (price_input.text.isEmpty || name_input.text.isEmpty || description_input.text.isEmpty) {
                                                    ScaffoldMessenger.of(context).showSnackBar(SnackBar(
                                                        content: Text('Please fill all the fields')));
                                                    
                                                  }
                                                else if (price_input.text.contains('abcdefghijklmnopqrstuvwxyz') || price_input.text.contains('ABCDEFGHIJKLMNOPQRSTUVWXYZ') || price_input.text.contains('!@#%^&*()_+')) {
                                                    ScaffoldMessenger.of(context).showSnackBar(SnackBar(
                                                        content: Text('Price should be a number')));
                                                    
                                                  }else{
                                                    final productBloc = BlocProvider.of<UpdateBloc>(context);
                                                    productBloc.add(UpdateProductEvent(product: updateModel));
                                                  }
                                      });
                                },
                              ),
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
                )
              ],
            ),
          );
        },
      ),
      // body: Text()
    );
  }
}
