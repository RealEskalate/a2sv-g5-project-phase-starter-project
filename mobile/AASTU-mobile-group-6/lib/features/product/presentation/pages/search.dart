import 'package:ecommerce_app_ca_tdd/features/product/presentation/bloc/home_bloc.dart';
import 'package:ecommerce_app_ca_tdd/features/product/presentation/bloc/home_event.dart';
import 'package:ecommerce_app_ca_tdd/features/user_auth/presentation/bloc/get_user/get_user_bloc.dart';
import 'package:ecommerce_app_ca_tdd/features/user_auth/presentation/bloc/get_user/get_user_event.dart';
import "package:flutter/material.dart";
import 'package:ecommerce_app_ca_tdd/extra/overflow_card.dart';
import 'package:flutter_bloc/flutter_bloc.dart';
import 'package:provider/provider.dart';
import 'package:ecommerce_app_ca_tdd/extra/search_func.dart';
import 'package:ecommerce_app_ca_tdd/features/product/presentation/pages/home.dart';
import 'package:ecommerce_app_ca_tdd/extra/resusetext.dart';
import 'package:ecommerce_app_ca_tdd/features/product/presentation/pages/details.dart';
import 'package:ecommerce_app_ca_tdd/models/product.dart';
import 'package:ecommerce_app_ca_tdd/models/ext_product.dart';

import '../bloc/home_state.dart';


class searchPage extends StatefulWidget {
  const searchPage({super.key});

  @override
  State<searchPage> createState() => _searchPageState();
}

class _searchPageState extends State<searchPage> {
      Future<void> _refresh() {
      context.read<HomeBloc>().add(GetProductsEvent());
      
    
  
  return  Future.delayed(Duration(seconds: 3));}
    
  @override
  Widget build(BuildContext context) {
    return Scaffold(
        appBar: AppBar(
          automaticallyImplyLeading: false,
          title: Row(
            mainAxisAlignment: MainAxisAlignment.spaceBetween,
            children: [
              IconButton(onPressed: () { Navigator.pushNamed(context,'/home');}, icon: Icon(Icons.arrow_back_ios_new,color: Color.fromARGB(255, 63, 81, 243),size: 20,)),
              const Center(
                child: Text("Search  Product"),
              ),
              const SizedBox(
                height: 60,
                width: 60,
              )
            ],
          ),
        ),


        // Body Starts Here
        body: SingleChildScrollView(
          child: Container(
            margin: EdgeInsets.only(left: 32,right: 24),
            child: Column(
              children: [
                Row(
                  // mainAxisAlignment: MainAxisAlignment.spaceBetween,
                  children: [
                    GestureDetector(
                      onTap: (){SearchFunc();},
                      child: SizedBox(
                        width: 300,
                        height: 48,
                        child: Container(
                          decoration: BoxDecoration(
                            border: Border.all(color: Color.fromRGBO(217, 217, 217, 1)),
                            borderRadius: BorderRadius.all(Radius.circular(8))
                            ),
                          child: const Expanded(
              
                            child: TextField(
                      
                              decoration: InputDecoration(
                                suffixIcon: Icon(Icons.arrow_forward),
                                border: InputBorder.none,
                                hintText: "  Leather",
                                
                              ), 
                            ),
                          ),
                        ),
                      ) 
                    ),
                    SizedBox(
                      width: 7,
                    ),
                    Container(

                      decoration: BoxDecoration(
                        color: Color.fromRGBO(63, 81, 243, 1),
                        
                      ),
                      child: SizedBox(
                        height: 48,
                        width: 48,
                        
                          child: Container(
                            decoration: BoxDecoration(
                              border: Border.all(width: 4,color: Color.fromRGBO(63, 81, 243, 1)),
                              borderRadius: BorderRadius.all(Radius.circular(8))
                            ),
                            child: IconButton(
                              onPressed: (){
                                showModalBottomSheet
                                (context: context, 
                                builder: (BuildContext context){
                                  return const SizedBox(
                                      height: 338,
                                      child:  about_product());
                                            
                                });
                              }, 
                              icon: Icon(Icons.filter_list,color: Colors.white,))),
                        ),
                    ),
                    
                
                  ],),
                  SizedBox(height:31),
                   BlocBuilder<HomeBloc, HomeState>(
                  builder: (context, state) {
                    if (state is HomeLoading) {
                      return Center(
                        child: CircularProgressIndicator());
                    } else if (state is HomeFailure) {
                      return SnackBar(
                        content: Text(state.message),);
                    }else if (state is HomeLoaded) {
                      return Expanded(
                      child: SizedBox(
                        child: SingleChildScrollView(
                          child: SizedBox(
                            height: MediaQuery.of(context).size.height * 0.8,
                            child: RefreshIndicator(
                              onRefresh: _refresh,
                              child: ListView.builder(
                                itemCount: state.products.length,
                                itemBuilder: (context, index) {
                                  return GestureDetector(
                                    onTap: () {
                                    Navigator.pushNamed(context, '/detail',arguments: state.products[index]);},
                                    child: OverflowCard(product: state.products[index],)
                                  );
                                },
                              ),
                            ),
                          ),
                        ),
                      ),
                                      );
                }
                return Container(); // Add a return statement at the end
              },
            ),
          //         
                    
              ],
            ),
          ),
        ) ,
    );
  }
}