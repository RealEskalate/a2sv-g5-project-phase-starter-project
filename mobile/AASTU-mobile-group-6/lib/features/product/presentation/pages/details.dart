// ignore_for_file: prefer_const_constructors, prefer_const_literals_to_create_immutables
import 'package:ecommerce_app_ca_tdd/features/product/data/models/product_models.dart';
import 'package:ecommerce_app_ca_tdd/features/product/domain/entities/product_entity.dart';
import 'package:ecommerce_app_ca_tdd/features/product/presentation/bloc/detail/detail_bloc.dart';
import 'package:ecommerce_app_ca_tdd/features/product/presentation/bloc/detail/detail_event.dart';
import 'package:ecommerce_app_ca_tdd/features/product/presentation/bloc/detail/detail_state.dart';
import 'package:ecommerce_app_ca_tdd/features/product/presentation/bloc/home_bloc.dart';
import 'package:ecommerce_app_ca_tdd/features/product/presentation/bloc/home_event.dart';
import 'package:ecommerce_app_ca_tdd/locator.dart';
import 'package:flutter/material.dart';
import 'package:flutter/widgets.dart';
import 'package:ecommerce_app_ca_tdd/features/product/presentation/pages/add_update.dart';
import 'package:ecommerce_app_ca_tdd/features/product/presentation/pages/update.dart';
import 'package:flutter_bloc/flutter_bloc.dart';
import 'package:google_fonts/google_fonts.dart';
import 'package:ecommerce_app_ca_tdd/extra/resusetext.dart';
import 'package:ecommerce_app_ca_tdd/models/product.dart';
import 'package:ecommerce_app_ca_tdd/models/ext_product.dart';

class DetailsPage extends StatefulWidget {
  final ProductModel item;
  const DetailsPage({required this.item, super.key});

  @override
  _DetailsState createState() => _DetailsState();
}

class _DetailsState extends State<DetailsPage> {
  var list = List<int>.generate(7, (i) => i + 1 + 38);
  int selected = 0;
  @override
  Widget build(BuildContext context) {
    var on = (){context.read<DetailBloc>().add(DeleteProductEvent(widget.item.id));};
    return BlocProvider(
      create: (context) => sl.get<DetailBloc>(),
      child: Scaffold(
            backgroundColor: Colors.white,
            floatingActionButtonLocation: FloatingActionButtonLocation.miniStartTop,
            floatingActionButton: Container(
              margin: EdgeInsets.only(left: 20, top: 20),
              height: 40,
              width: 40,
              child: FloatingActionButton(
                  shape: CircleBorder(),
                  backgroundColor: Colors.white,
                  child: Center(
                      child: Icon(
                    Icons.arrow_back_ios_new,
                    color: Color.fromARGB(255, 63, 81, 243),
                    size: 20,
                  )),
                  onPressed: () {
                    Navigator.pop(context);
                  }),
            ),
            body: SafeArea(
              child: BlocConsumer<DetailBloc, DetailState>(
                listener: (context, state) {
                    if(state is DetailLoaded){
                      ScaffoldMessenger.of(context).showSnackBar(
                      SnackBar(
                        backgroundColor: Colors.teal[100],
                    content: Row(
                      // mainAxisAlignment: MainAxisAlignment.spaceEvenly,
                      children: [
                        SizedBox(width: 25,),
                        Icon(Icons.thumb_up_rounded,color: Colors.yellow,),
                        SizedBox(width: 60,),
                        Text("Succesfully Deleted !!!!",style: TextStyle(color: Colors.white),),
                      ],
                    ),
                        
                      ));
                      context.read<HomeBloc>().add(GetProductsEvent());
                  
                      Navigator.pushNamed(context,'/home');
                    }
                },
                builder: (context, state) {
                  return SingleChildScrollView(
                    child: Column(
                      crossAxisAlignment: CrossAxisAlignment.start,
                      mainAxisAlignment: MainAxisAlignment.start,
                      children: [
                        Image(
                          height: 286,
                          width: 430,
                          fit: BoxFit.cover,
                          image: NetworkImage(widget.item.imagePath),
                        ),
                        Column(
                          children: [
                            Row(
                              mainAxisAlignment: MainAxisAlignment.spaceBetween,
                              // crossAxisAlignment: CrossAxisAlignment.,
                              children: [
                                Container(
                                  margin: EdgeInsets.only(top: 21),
                                  child: Column(
                                    crossAxisAlignment: CrossAxisAlignment.start,
                                    children: [
                                      Container(
                                          margin: EdgeInsets.only(left: 32),
                                          child: Text(
                                            "Men" + "'" + "s" + "Shoes",
                                            style: GoogleFonts.poppins(
                                                fontSize: 16,
                                                fontWeight: FontWeight.w400,
                                                color: Color.fromARGB(
                                                    255, 170, 170, 170)),
                                          )),
                                      Container(
                                          margin: EdgeInsets.only(top: 16, left: 32),
                                          child: Text(
                                            widget.item.name,
                                            style: GoogleFonts.poppins(
                                                fontSize: 24,
                                                fontWeight: FontWeight.w600),
                                          )),
                                    ],
                                  ),
                                ),
                                Column(
                                  crossAxisAlignment: CrossAxisAlignment.end,
                                  children: [
                                    Container(
                                      margin: EdgeInsets.only(right: 32),
                                      child: Row(
                                        children: [
                                          Icon(
                                            Icons.star,
                                            size: 26,
                                            color: Color.fromARGB(255, 255, 215, 0),
                                          ),
                                          Text("(4.0)",
                                              style: GoogleFonts.sora(
                                                  fontSize: 16,
                                                  fontWeight: FontWeight.w400,
                                                  color: Color.fromARGB(
                                                      255, 170, 170, 170))),
                                        ],
                                      ),
                                    ),
                                    Container(
                                        margin: EdgeInsets.only(top: 23, right: 32),
                                        child: Text(
                                          "\$" + widget.item.price.toString(),
                                          style: GoogleFonts.sora(
                                              fontSize: 16,
                                              fontWeight: FontWeight.w500),
                                        ))
                                  ],
                                ),
                              ],
                            ),
                            SizedBox(
                              height: 10,
                            ),
                            Container(
                              margin: EdgeInsets.only(right: 298, left: 10),
                              child: Text(
                                "Sizes:",
                                style: GoogleFonts.poppins(
                                    fontWeight: FontWeight.w500, fontSize: 20),
                              ),
                            ),
                            // SIZES PART
                            Container(
                              margin: EdgeInsets.only(left: 32),
                              child: Row(
                                children: [
                                  SizedBox(
                                    height: 60,
                                    width: 378,
                                    child: ListView.builder(
                                      itemCount: 7,
                                      scrollDirection: Axis.horizontal,
                                      itemBuilder: (context, index) {
                                        return GestureDetector(
                                          onTap: () {
                                            setState(() {
                                              selected = index;
                                            });
                                          },
                                          child: Card(
                                            color: selected == index
                                                ? const Color(0xff3F51F3)
                                                : Colors.white,
                                            child: SizedBox(
                                                height: 60,
                                                width: 60,
                                                child: Center(
                                                    child: reusableText(
                                                        "${list[index]}",
                                                        FontWeight.w500,
                                                        17,
                                                        selected == index
                                                            ? Colors.white
                                                            : Colors.black))),
                                          ),
                                        );
                                      },
                                    ),
                                  ),
                                ],
                              ),
                            ), // END OF SIZES Part
                  
                            Row(
                              children: [
                                Expanded(
                                    child: Padding(
                                  padding: const EdgeInsets.only(
                                      left: 32, right: 32, top: 15, bottom: 5),
                                  child: Center(
                                    child: Text(
                                      widget.item.description,
                                      style: GoogleFonts.poppins(
                                          color: Color.fromRGBO(102, 102, 102, 1),
                                          fontSize: 14,
                                          fontWeight: FontWeight.w500),
                                    ),
                                  ),
                                ))
                              ],
                            ),
                  
                            Padding(
                              padding:
                                  const EdgeInsets.only(left: 32, right: 32, top: 20),
                              child: Row(
                                mainAxisAlignment: MainAxisAlignment.spaceBetween,
                                children: [
                                  
                                  SizedBox(
                                    width: 152,
                                    height: 50,
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
                                          var on_press = (){context.read<DetailBloc>().add(DeleteProductEvent(widget.item.id));};
                                          showDialog(
                                            context: context, 
                                            builder: (context)=> AlertDialog(
                                              title: Text("Are you sure you want to delete this product?",style: GoogleFonts.poppins(fontSize: 15),),
                                              actions: [
                                                TextButton(onPressed: (){
                                                  Navigator.pop(context);
                                                }, child: Text("Cancel")),
                                                TextButton(style: ButtonStyle(backgroundColor: MaterialStatePropertyAll<Color>(Colors.red)),
                                                  onPressed: on_press(), child: Text("Delete"))
                                              ],
                                            )
                                          );
                  
                                      
                        
                                          // Navigator.pushNamed(context, '/home');
                                        },
                                        child: Text("Delete",
                                            style: GoogleFonts.poppins(
                                                fontSize: 14,
                                                fontWeight: FontWeight.w500))),
                                  ),
                                  SizedBox(
                                    width: 152,
                                    height: 50,
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
                                          Navigator.pushNamed(context, '/update',
                                              arguments: widget.item);
                                        },
                                        child: Text(
                                          "Update",
                                          style: GoogleFonts.poppins(
                                              fontSize: 14,
                                              fontWeight: FontWeight.w500),
                                        )),
                                  ),
                                ],
                              ),
                            ) // End of Paragraph
                          ],
                        )
                      ],
                    ),
                  );
                },
              ),
            )),
    );
  }
}

//End of Details Page
